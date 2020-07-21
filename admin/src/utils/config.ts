let baseURL = 'http://go.abadboy.cn/api/v1';
const isDebug = process.env.NODE_ENV === 'development';
if (isDebug) {
  baseURL = 'http://127.0.0.1:8080/api/v1';
}
export default {
  previewUrl: 'http://go.abadboy.cn',
  baseURL,
  isDebug,
};
