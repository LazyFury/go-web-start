import { AxiosResponse } from 'axios';
import { useEffect, useState } from 'react';

export interface listResult {
  list: Array<any>;
  page: number;
  size: number;
  page_count: number;
  total: number;
}

const defaultData: listResult = {
  list: [],
  page: 0,
  size: 0,
  page_count: 0,
  total: 0,
};

/**
 * @获取数据列表hooks
 * @param api
 * @see https://xxx.com
 */
export function useDataList(
  api: (page: number) => Promise<AxiosResponse<listResult>>,
  autoLoad: boolean | undefined = true,
): {
  readonly data: listResult;
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
      setData(res.data);
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
