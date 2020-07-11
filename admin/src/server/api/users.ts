import { http } from '../request';

export const users = {
  total: (params: object) => http.get(`/users-actions/count`, { params }),
};
