window.ws = null;

var app = new Vue({
  el: '#app',
  data: { val: '', ws: new SocketClient() },
  created() {
    this.start();
    this.ws.useOnMessage((e) => {
      if (!e.is_self) return;
      this.$nextTick(function () {
        document.querySelector('#output').scrollTop = 999 * 999 * 999;
      });
    });
  },
  methods: {
    start() {
      if (this.ws.conn != null) {
        if (confirm('链接似乎已建立，是否重连')) {
          this.ws.conn.close();
          this.ws.conn = null;
          setTimeout(() => {
            this.ws.connect();
          }, 300);
        }
      } else {
        this.ws.connect();
      }
    },
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
