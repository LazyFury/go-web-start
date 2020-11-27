import { http } from '../request';
/**
 * install 快速注册路由
 * @param name
 */
export function install(name: string) {
  return {
    list: (params: object) => http.get('/' + name, { params }),
    del: (id: number) => http.delete(`/${name}/${id}`),
    detail: (id: number) => http.get(`/${name}/${id}`),
    add: (data: object) => http.post(`/${name}`, data),
    update: (id: any, data: object) => http.patch(`/${name}/${id}`, data),
    total: (params: object) => http.get(`/${name}-actions/count`, { params }),
  };
}
