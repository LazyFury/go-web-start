import { http } from '../request';
import { install } from './easy_install';

export const posts = {
  ...install('articles'),
};

export const postCates = {
  ...install('article-cates'),
  list: () => http.get(`/article-cates-all`),
};

export const postRec = {
  ...install('article-recs'),
  list: () => http.get('/article-recs-all'),
};

export const postTags = {
  ...install('article-tags'),
  list: () => http.get('/article-tags-all'),
};
