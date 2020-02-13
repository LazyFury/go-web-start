<template>
  <a-drawer
    title="添加API参数"
    placement="left"
    width="500"
    :closable="true"
    @close="onClose"
    :visible="visible"
  >
    <a-form :form="form" @submit="handleSubmit">
      <a-form-item label="参数名称" :label-col="{ span: 5 }" :wrapper-col="{ span: 12 }">
        <a-input v-decorator="['name', { rules: [{ required: true, message: '请输入提示名称!' }] }]" />
      </a-form-item>
      <a-form-item label="Key" :label-col="{ span: 5 }" :wrapper-col="{ span: 12 }">
        <a-input v-decorator="['key', { rules: [{ required: true, message: '请选择key名称!' }] }]" />
      </a-form-item>
      <a-form-item label="Value" :label-col="{ span: 5 }" :wrapper-col="{ span: 12 }">
        <a-input v-decorator="['value', { rules: [{ required: true, message: '请选择默认值!' }] }]" />
      </a-form-item>
      <a-form-item label="类型" :label-col="{ span: 5 }" :wrapper-col="{ span: 12 }">
        <a-select
          v-decorator="['type', { rules: [{ required: true, message: '请选择默认值!' }],initialValue: 'string' }]"
          style="width: 120px"
        >
          <a-select-option value="string">string</a-select-option>
          <a-select-option value="number">number</a-select-option>
          <!-- <a-select-option value="disabled" disabled>Disabled</a-select-option> -->
        </a-select>
      </a-form-item>

      <a-form-item label="提示内容" :label-col="{ span: 5 }" :wrapper-col="{ span: 12 }">
        <a-input v-decorator="['tips', { rules: [{ required: false, message: '请选择默认值!' }] }]" />
      </a-form-item>
      <a-form-item :wrapper-col="{ span: 12, offset: 5 }">
        <a-button type="primary" html-type="submit">保存</a-button>
        <!-- <a-button type @click="reg">注册</a-button> -->
      </a-form-item>
    </a-form>
  </a-drawer>
</template>

<script>
export default {
  data() {
    return {
      formLayout: "horizontal",
      form: this.$form.createForm(this, { name: "coordinated" }),
      visible: false
      // defaultData: {}
    };
  },
  methods: {
    show(config) {
      this.visible = true;
      if (config) {
        // this.defaultData = config;
        console.log(this.form);
        this.$nextTick(() => {
          this.form.setFieldsValue(config);
        });
      }
    },
    onClose() {
      this.visible = false;
    },
    handleSubmit(e) {
      e.preventDefault();
      this.form.validateFields((err, values) => {
        if (!err) {
          console.log("Received values of form: ", values);
          this.$emit("save", values);
          this.form.resetFields();
          // alert(this.visible);
          this.visible = false;
        }
      });
    }
  }
};
</script>