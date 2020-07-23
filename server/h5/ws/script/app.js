window.ws = null;

var app = new Vue({
  el: '#app',
  data: { val: '', ws: new SocketClient() },
  created() {
    this.ws.connect();
  },
  methods: {
    send() {
      this.ws.send(this.val);
    },
    confirm() {
      this.send();
    },
  },
});

function keyDown(e) {
  console.log(e);
}
