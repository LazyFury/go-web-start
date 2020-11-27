import { http } from '../request';
import { install } from './easy_install';

export const posts = {
  ...install('articles'),
};

export const postCates = {
  ...install('post-cates'),
  list: () => http.get(`/post-cates`),
};

export const postRec = {
  ...install('article-recs'),
  list: () => http.get('/article-recs-all'),
};

export const postTags = {
  ...install('post-tags'),
  list: () => http.get('/post-tags'),
};
