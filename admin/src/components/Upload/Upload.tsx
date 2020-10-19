import { http } from '@/server/request';
import { CloudUploadOutlined } from '@ant-design/icons';
import { Row } from 'antd';
import Button from 'antd/es/button/button';
import React, { useState } from 'react';

interface uploadValues {
  url: string;
  file: uploadFile;
  status: 1 | 2 | 3 | 4;
}

interface uploadProps {
  value?: uploadValues[];
  onchange?: (value: uploadValues[]) => void;
}

class uploadFile {
  event: Event;
  el: HTMLFormElement;
  files: File[];
  constructor(e: Event) {
    this.event = e;
    this.el = this.event?.path[0] || null;
    this.files = this.el?.files;
  }

  get formData() {
    var formData = new FormData();
    this.files?.forEach(x => {
      formData.append('file', x);
    });
    return formData;
  }
}
export const Uploader: React.FC<uploadProps> = ({ value, onchange }) => {
  let [list, setList] = useState(value);
  const addFile = (file: uploadValues) => {
    setList([...(list || []), file]);
    onchange && list && onchange(list);
  };
  const chooseFile = () => {
    let input = document.createElement('input');
    input.type = 'file';
    input.multiple = false;
    input.onchange = e => {
      let file = new uploadFile(e);
      console.log(file);
      http
        .post('/upload-img', file.formData, {
          headers: { 'Content-Type': 'multipart/form-data;' },
          transformRequest: [data => data],
          onUploadProgress: e => {
            let complete = (((e.loaded / e.total) * 100) | 0) + '%';
            console.log(complete);
          },
        })
        .then(res => {
          addFile({ url: res.data, file: file, status: 1 });
        });
      input.remove();
    };
    input.click();
  };
  return (
    <Row>
      <Button onClick={chooseFile}>
        <CloudUploadOutlined />
        upload file
      </Button>

      {JSON.stringify(list)}
    </Row>
  );
};
