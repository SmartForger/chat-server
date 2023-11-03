const lib = ChatLib({
    server: '',
});

async function init() {
    await lib.loadLocalData();

    const isLoggedIn = await lib.login();
    if (isLoggedIn) {
        initChat();
    } else {
        showLoginScreen();
    }
}

function showLoginScreen() {
    const appContainer = document.getElementById('app');
    appContainer.innerHTML = '';

    const formEl = document.createElement('form');
    formEl.id = 'loginform';
    formEl.innerHTML = `
<div>
    <label for="username">Username</label>
    <input id="username" class="input" placeholder="Enter Username">
</div>
<div>
    <label for="password">Password</label>
    <input id="password" class="input" type="password" placeholder="Enter Password">
</div>
<input class="primary-btn" type="submit" value="Log In">
`;

    formEl.addEventListener('submit', async (ev) => {
        if (ev) {
            ev.preventDefault();
        }
    
        const isLoggedIn = await lib.login({
            Username: document.getElementById('username').value,
            Password: document.getElementById('password').value,
        });
    
        if (isLoggedIn) {
            initChat();
        }
    });

    appContainer.appendChild(formEl);
}

function initChat() {
    showChatScreen();
    initSocket();
}

function initSocket() {
    const socket = io('');

    socket.on('joined', function(msg){
        console.log('Joined room', msg);
    });

    socket.on('connect', async () => {
        const client = lib.getClient();
        const msg = await lib.getSocketMessage({ Username: client.Username });
        socket.emit('join', msg);
    });

    socket.on('custom_error', (msg) => {
        console.error('socket error', msg);
        if (msg === 'join_error') {
            socket.close();
            window.location.reload();
        }
    });
}

function showChatScreen() {
    const appContainer = document.getElementById('app');
    appContainer.innerHTML = '';

    const messages = document.createElement('div');
    messages.id = 'messages';
    appContainer.appendChild(messages);

    const toolbar = document.createElement('div');
    toolbar.id = 'toolbar';
    toolbar.innerHTML = `
        <textarea id="input" class="input"></textarea>
        <button id="send" class="primary-btn">Send</button>
    `;

    appContainer.appendChild(toolbar);
}

function addSystemMessage(msg) {
    const container = document.getElementById('messages');

    const msgEl = document.createElement('div');
    msgEl.className = 'message system';
    msgEl.innerText = msg;

    container.appendChild(msgEl);
}

function addMessage(msg) {
    const container = document.getElementById('messages');

    const now = new Date();
    const h = now.getHours();
    const timeStr = `${h % 12}:${now.getMinutes()} ${h >= 12 ? 'pm' : 'am'}`;

    const msgEl = document.createElement('div');
    msgEl.className = 'message';
    msgEl.innerHTML = `
    <div class="time">${timeStr}</div>
    <div class="text">${msg}</div>
    `;

    container.appendChild(msgEl);
}

init();
