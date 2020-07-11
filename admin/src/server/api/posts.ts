import { http } from '../request';

export const posts = {
  list: (params: object) => http.get('/posts', { params }),
  del: (id: number) => http.delete(`/posts/${id}`),
  detail: (id: number) => http.get(`/posts/${id}`),
  add: (data: object) => http.post(`/posts`, data),
  update: (id: any, data: object) => http.put(`/posts/${id}`, data),
  total: (params: object) => http.get(`/posts-actions/count`, { params }),
};

export const postCates = {
  list: () => http.get(`/post-cates`),
};
