<template>
  <a-form :form="form" @submit="handleSubmit">
    <a-form-item label="参数名" :label-col="{ span: 5 }" :wrapper-col="{ span: 12 }">
      <a-input v-decorator="['key', { rules: [{ required: true, message: '请选择key名称!' }] }]" />
    </a-form-item>
    <a-form-item label="默认值" :label-col="{ span: 5 }" :wrapper-col="{ span: 12 }">
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
    <a-form-item label="提示名称" :label-col="{ span: 5 }" :wrapper-col="{ span: 12 }">
      <a-input v-decorator="['name', { rules: [{ required: true, message: '请输入提示名称!' }] }]" />
    </a-form-item>
    <a-form-item :wrapper-col="{ span: 12, offset: 5 }">
      <a-button type="primary" html-type="submit">保存</a-button>
      <!-- <a-button type @click="reg">注册</a-button> -->
    </a-form-item>
  </a-form>
</template>

<script>
export default {
  data() {
    return {
      formLayout: "horizontal",
      form: this.$form.createForm(this, { name: "coordinated" })
    };
  },
  methods: {
    handleSubmit(e) {
      e.preventDefault();
      this.form.validateFields((err, values) => {
        if (!err) {
          console.log("Received values of form: ", values);
          this.$emit("save", values);
          this.form.resetFields();
        }
      });
    },
    handleSelectChange(value) {
      console.log(value);
      this.form.setFieldsValue({
        note: `Hi, ${value === "male" ? "man" : "lady"}!`
      });
    }
  }
};
</script>