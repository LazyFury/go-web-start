class SocketClient {
  conn = null;
  userId = '';
  count = 0;
  onlineUser = {};

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
    uri += '/api/v1/ws';
    this.conn = new WebSocket(uri);
    this.conn.onopen = () => this.onOpen();
    this.conn.onmessage = (e) => this.onMessage(e);
    this.conn.onclose = (e) => this.onClose(e);
  }

  onOpen() {
    console.log('Connected');
    this.send('join'); //服务端第一次收到消息之后创建用户，发送用户加入通知
  }

  onMessage(evt) {
    let result = {};
    try {
      result = JSON.parse(evt.data);
    } catch (err) {
      throw err;
    }
    let id = result.ID;
    console.log(result, id);
    id && window.localStorage.setItem('userId', id);
    this.userId = id;

    console.log(this.messages);
    this.messages.push(result);
    // update
    this.count = result.count;
    this.onlineUser = result.OnLineUser;
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
    let id = this.userId;
    if (!id) {
    }
    if (this.conn) this.conn.send(JSON.stringify({ msg: val, action, id }));
  }
}
