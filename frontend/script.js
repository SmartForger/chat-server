const lib = ChatLib({
    server: '',
});

async function init() {
    const form = document.getElementById('loginform');

    await lib.loadLocalData();

    const isLoggedIn = await lib.login();
    if (isLoggedIn) {
        form.remove();
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
    }
}

function addMessage(msg) {
    const li = document.createElement('li');
    li.innerText = msg;
    document.getElementById('messages').appendChild(li);
}

function initSocket() {
    var socket = io('');
    var s2 = io("/chat");

    socket.on('reply', function(msg){
        addMessage(msg);
    });

    const form1 = document.getElementById('inputform');
    const msgEl = document.getElementById('m');

    form1.addEventListener('submit', (ev) => {
        ev.preventDefault();

        s2.emit('msg', msgEl.value, function(data){
            addMessage('ACK CALLBACK: ' + data);
        });

        socket.emit('notice', msgEl.value);

        msgEl.value = '';
    });
}

init();
