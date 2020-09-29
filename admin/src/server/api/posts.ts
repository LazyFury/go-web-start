import { http } from '../request';
import { install } from './easy_install';

export const posts = {
  ...install('posts'),
};

export const postCates = {
  ...install('post-cates'),
  list: () => http.get(`/post-cates`),
};

export const postRec = {
  ...install('post-rec'),
  list: () => http.get('/post-rec'),
};

export const postTags = {
  ...install('post-tags'),
  list: () => http.get('/post-tags'),
};
