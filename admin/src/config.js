let baseURL = "http://127.0.0.1:8080"

if (process.env.NODE_ENV !== 'development') {
  baseURL = ''
}
export default {
  baseURL
}
