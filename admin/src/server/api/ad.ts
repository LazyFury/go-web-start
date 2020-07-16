import { http } from '../request';
import { install } from './easy_install';

export const ads = {
  ...install('ads'),
  list: () => http.get('/ads'),
};

export const adEvents = {
  ...install('ad-events'),
  list: () => http.get('/ads'),
};

export const adGroup = { ...install('ad-groups') };
