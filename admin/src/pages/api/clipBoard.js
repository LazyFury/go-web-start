import clipboard from 'clipboard'
import { message } from 'ant-design-vue'

export const clipBoardText = (text = "") => {
    initClipBoard();
    console.log(text)
    let button = document.createElement('button')
    button.setAttribute('data-clipboard-action', 'copy')
    button.setAttribute('data-clipboard-text', text)
    button.className = 'btn'
    document.body.appendChild(button)
    button.click()
    button.remove()
}
window.clipBoardText = clipBoardText


function initClipBoard() {
    if (window.clipboard_) {
        window.clipboard_.destroy()
    }
    window.clipboard_ = new clipboard('.btn');
    // this.$alert("onLoad")
    window.clipboard_.on('success', function (e) {
        console.log(e);
        message.success("复制成功");
    });
    window.clipboard_.on('error', function (e) {
        console.log(e);
        message.error("复制失败");
    });
}
