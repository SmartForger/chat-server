const lib = ChatLib({
    server: '',
});

async function init() {
    const form = document.getElementById('loginform');

    await lib.loadLocalData();

    const isLoggedIn = await lib.login();
    if (isLoggedIn) {
        form.remove();
        initSocket();
    } else {
        form.addEventListener('submit', login);
    }
}

async function login(ev) {
    if (ev) {
        ev.preventDefault();
    }

    const isLoggedIn = await lib.login({
        Username: document.getElementById('username').value,
        Password: document.getElementById('password').value,
    });

    if (isLoggedIn) {
        const form = document.getElementById('loginform');
        form.remove();
        initSocket();
    }
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

init();
