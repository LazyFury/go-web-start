import { List, Result } from '@/server/interface';
import { AxiosResponse } from 'axios';
import { useEffect, useState } from 'react';
let defaultData: List<any> = {
  list: [],
  page: 1,
  page_count: 1,
  total: 0,
  size: 10,
};
/**
 * @获取数据列表hooks
 * @param api
 * @see https://xxx.com
 */
export function useDataList(
  api: (page: number) => Promise<AxiosResponse<Result<List<any>>>>,
  autoLoad: boolean | undefined = true,
): {
  readonly data: List<any>;
  load: (p?: number | undefined) => Promise<void>;
  reset: () => Promise<void>;
  reload: () => Promise<void>;
  loading: boolean;
} {
  let [page, setPage] = useState(1);
  let [data, setData] = useState(defaultData);
  let [loading, setLoading] = useState(true);

  const load = async (p?: number) => {
    setLoading(true);
    const res = await api(p || page);
    if (res) {
      console.log(res);
      setData(res.data.data);
      setPage(p || page++);
    }
    setLoading(false);
  };

  const reset = () => load(1); //从第一页开始
  const reload = () => load(page); //刷新当前页

  useEffect(() => {
    if (autoLoad) load();
  }, []);
  return { data, load, reset, reload, loading };
}
