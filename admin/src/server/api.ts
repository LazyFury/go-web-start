import { http } from './request';

let postAPI = '/posts';
export const posts = {
  list: (params: object) => http.get(postAPI, { params }),
  del: (id: number) => http.delete(`${postAPI}/${id}`),
  detail: (id: number) => http.get(`${postAPI}/${id}`),
  add: (data: object) => http.post(postAPI, data),
  update: (data: object) => http.put(postAPI, data),
};
