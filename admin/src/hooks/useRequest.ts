import { AxiosResponse } from 'axios';
import { useEffect, useState } from 'react';

const useRequest = (
  api: () => Promise<AxiosResponse<any>>,
  autoLoad: boolean | undefined = false,
): { readonly data: any; load: () => Promise<any> } => {
  let [data, setData] = useState({});

  const load = () =>
    api().then(res => {
      if (res) {
        setData(res.data);
      }
    });

  useEffect(() => {
    if (autoLoad) {
      load();
    }
  }, []);

  return { data, load };
};
export default useRequest;
