import config from '@/utils/config';
import { CloudUploadOutlined } from '@ant-design/icons';
import { message, Upload } from 'antd';
import Button from 'antd/es/button/button';
import React, { useState } from 'react';

interface uploadProps {
  value?: string[] | string;
  onChange?: (value: string[] | string) => void;
  multiple: boolean;
}

export const Uploader: React.FC<uploadProps> = ({
  value,
  onChange,
  multiple,
}) => {
  let [fileArr, setFileArr] = useState<any[]>([]);

  const updateValue = (arr: any[]) => {
    if (onChange) {
      if (!multiple) {
        onChange(arr[0]?.response?.data || '');
      } else {
        onChange(
          arr.map((x: { response: { data: any } }) => x?.response?.data),
        );
      }
    }
  };
  const handleChange = (info: any) => {
    let _fileList = [...info.fileList];
    _fileList = _fileList.slice(-1);
    setFileArr(_fileList);

    updateValue(_fileList);

    if (info.file.status !== 'uploading') {
      console.log(info.file, info.fileList);
    }
    if (info.file.status === 'done') {
      message.success(`${info.file.name} 上传成功`);
    } else if (info.file.status === 'error') {
      message.error(
        `${info.file.name} ${info.file?.response?.msg || '上传失败'} `,
      );
    }
  };
  const props = {
    name: 'file',
    action: config.baseURL + '/upload-img',
    headers: {
      Authorization: window.localStorage.getItem('token') || '',
    },
    multiple: multiple,
    onChange: handleChange,
  };

  return (
    <Upload {...props} fileList={fileArr}>
      <Button>
        <CloudUploadOutlined />
        upload file
      </Button>
    </Upload>
  );
};
