/**
 * // 分割一个长数组至 n 个 len 长度的小数组
 * @param {Array} arr 需要转换的数组
 * @param {Number} len 切割的长度
 */
const split_array = (arr, len) => {
    var a_len = arr.length;
    var result = [];
    for (var i = 0; i < a_len; i += len) {
        result.push(arr.slice(i, i + len));
    }
    return result;
}

export default split_array