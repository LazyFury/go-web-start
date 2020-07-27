class SocketClient {
  conn = null;
  userId = '';
  count = 0;
  onlineUser = {};
  userName = '';
  useOnMessageFuncs = [];
  timer = null;

  constructor() {
    this.messages = [];
  }

  connect() {
    var loc = window.location;
    var uri = 'ws:';
    if (loc.protocol === 'https:') {
      uri = 'wss:';
    }
    uri += '//' + loc.host;
    uri += '/api/v1/ws?token=' + window.localStorage.getItem('userId');
    this.conn = new WebSocket(uri);
    this.conn.onopen = () => this.onOpen();
    this.conn.onmessage = (e) => this.onMessage(e);
    this.conn.onclose = (e) => this.onClose(e);
  }

  onOpen() {
    console.log('Connected');
    this.send('_', 'join'); //服务端第一次收到消息之后创建用户，发送用户加入通知
    this.heartbeat();
  }

  heartbeat() {
    if (this.timer != null) {
      clearInterval(this.timer);
      this.timer = null;
    }
    this.timer = setInterval(() => {
      this.send('_', 'ping');
    }, 1000 * 10);
  }

  onMessage(evt) {
    let result = {};
    try {
      result = JSON.parse(evt.data);
    } catch (err) {
      throw err;
    }

    // console.log(this.messages);

    switch (result.action) {
      case 'allUser':
      case 'SystemNotify':
        this.messages.push(result);
        break;
      case 'regUser':
        let id = result.from.id;
        console.log(result, id);
        id && window.localStorage.setItem('userId', id);
        this.userId = id;
        this.userName = result.from.name;
        break;
      case 'update':
        this.count = result.global.count;
        this.onlineUser = result.global.onlineUser;
        break;
    }

    this.useOnMessageFuncs.forEach((fn) => {
      fn.call(this, result);
    });
  }

  onClose(evt) {
    alert('链接已断开');
    console.error(evt);
  }

  // methods
  send(val, action = 'sendAll') {
    if (!val) {
      throw 'err 没有输入';
    }
    let id = window.localStorage.getItem('userId') || this.userId || '';
    if (!id) {
    }
    if (this.conn) this.conn.send(JSON.stringify({ msg: val, action, id }));
  }

  useOnMessage(fn) {
    if (fn instanceof Function) {
      this.useOnMessageFuncs.push(fn);
    } else {
      throw '不是合法的的方法';
    }
  }
}
