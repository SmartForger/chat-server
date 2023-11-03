async function init() {
    initSocket();
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
