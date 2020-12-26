import { http } from '../request';
import { install } from './easy_install';

export const ads = {
  ...install('ads'),
  // list: () => http.get('/ads'),
};

export const adEvents = {
  ...install('ad-events'),
  // override
  // list: () => http.get('/ad-events'),
  all: () => http.get('/ad-events-all'),
};

export const adGroups = {
  ...install('ad-groups'),
  // list: () => http.get('/ad-groups'),
  all: () => http.get('/ad-groups-all'),
};
