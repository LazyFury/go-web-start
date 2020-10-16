import config from '@/utils/config';
import { message } from 'antd';
import { UploadProps } from 'antd/lib/upload/interface';

export const defaultUploadProps: UploadProps<any> = {
  name: 'file',
  multiple: true,
  action: config.baseURL + '/upload-img',
  transformFile: file => {
    return file;
  },
  onChange: (info: any) => {
    const { status } = info.file;
    if (status !== 'uploading') {
      console.log(info.file, info.fileList);
    }
    if (status === 'done') {
      message.success(`${info.file.name} 上传成功`);
    } else if (status === 'error') {
      message.error(
        `${info.file.name} ${info?.file?.response?.msg ||
          '不是正确到文件类型'} `,
      );
    }
  },
};
