window.ws = null;

var app = new Vue({
  el: '#app',
  data: { val: '', ws: new SocketClient() },
  created() {
    this.ws.connect();
    this.ws.useOnMessage((e) => {
      if (!e.is_self) return;
      this.$nextTick(function () {
        document.querySelector('#output').scrollTop = 999 * 999 * 999;
      });
    });
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
