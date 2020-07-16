import { install } from './easy_install';

// 商品
export const products = {
  ...install('products'),
};
// 商品分类
export const productCates = {
  ...install('product-cates'),
};
