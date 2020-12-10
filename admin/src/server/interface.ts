export interface Result<T> {
  build_by: string;
  data: T;
  message: string;
  code: number;
}

export interface List<T> {
  list: Array<T>;
  page: number;
  page_count: number;
  size: number;
  total: number;
}
