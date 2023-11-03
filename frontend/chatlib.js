const ChatLib = ({ server }) => {
  const STORAGE_KEY_CLIENT = "client";
  const STORAGE_KEY_SECRET = "secret";
  const STORAGE_KEY_PUBLICKEY = "publickey";

  let publicKey = "";
  let client = {};

  async function getPublicKey() {
    if (publicKey) {
      return publicKey;
    }

    try {
      const res = await fetch(`${server}/_key`);
      const data = await res.json();

      publicKey = data.key;

      return data.key;
    } catch {
      return "";
    }
  }

  function generateNonce(publicKey) {
    const nonce = forge.random.getBytesSync(32);

    const rsapub = forge.pki.publicKeyFromPem(publicKey);
    const encryptedText = rsapub.encrypt(
      forge.util.encodeUtf8(btoa(nonce)),
      "RSA-OAEP",
      {
        md: forge.md.sha256.create(),
      }
    );

    return {
      nonce,
      encrypted: btoa(encryptedText),
    };
  }

  function encryptPayload(data, key) {
    const iv = forge.random.getBytesSync(32);
    const cipher = forge.cipher.createCipher("AES-GCM", key);
    cipher.start({ iv: iv });
    cipher.update(forge.util.createBuffer(data));
    cipher.finish();

    return btoa(iv + cipher.output.data + cipher.mode.tag.data);
  }

  function decryptPayload(data, key) {
    const encrypted = atob(data.slice(1, -1));

    const iv = encrypted.slice(0, 32);
    const tag = encrypted.slice(-16);
    const ciphertext = encrypted.slice(32, -16);
    const decipher = forge.cipher.createDecipher("AES-GCM", key);
    decipher.start({ iv: iv, tagLength: 128, tag });
    decipher.update(forge.util.createBuffer(ciphertext));
    decipher.finish();

    return decipher.output.data;
  }

  async function login(data) {
    if (client.Secret) {
      return true;
    }

    if (!data || !data.Username || !data.Password) {
      return false;
    }

    try {
      const {
        payload: { S: encryptedNonce, T: encryptedData },
        nonce,
      } = await getEncryptedPayload(data);

      const resp = await fetch(`${server}/client/login`, {
        method: "POST",
        body: encryptedData,
        headers: {
          sync_nonce: encryptedNonce,
        },
      });

      const responseData = await resp.text();
      if (!resp.ok) {
        throw {
          status: resp.status,
          data: responseData,
        };
      }

      client = JSON.parse(decryptPayload(responseData, nonce));

      localStorage.setItem(STORAGE_KEY_CLIENT, responseData);
      localStorage.setItem(STORAGE_KEY_SECRET, btoa(nonce));
      localStorage.setItem(STORAGE_KEY_PUBLICKEY, publicKey);

      return true;
    } catch (e) {
      console.error(e);
    }

    return false;
  }

  async function loadLocalData() {
    const pkey = localStorage.getItem(STORAGE_KEY_PUBLICKEY);
    const pkey1 = await getPublicKey();

    if (!pkey || pkey !== pkey1) {
      return;
    }

    const encryptedClient = localStorage.getItem(STORAGE_KEY_CLIENT);
    const secret = localStorage.getItem(STORAGE_KEY_SECRET);
    client = JSON.parse(decryptPayload(encryptedClient, atob(secret)));
  }

  async function getEncryptedPayload(data) {
    const publicKey = await getPublicKey();
    const { nonce, encrypted } = generateNonce(publicKey);
    const encryptedData = encryptPayload(JSON.stringify(data), nonce);

    return { payload: { S: encrypted, T: encryptedData }, nonce };
  }

  async function getSocketMessage(data) {
    try {
      const { payload } = await getEncryptedPayload(data);
      return JSON.stringify(payload);
    } catch {
      return "";
    }
  }

  function getClient() {
    return client;
  }

  return {
    login,
    loadLocalData,
    getClient,
    getSocketMessage,
  };
};
