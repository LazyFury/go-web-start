<template>
  <a-form :form="form" @submit="handleSubmit">
    <a-form-item label="用户名" :label-col="{ span: 5 }" :wrapper-col="{ span: 12 }">
      <a-input
        v-decorator="['username', { rules: [{ required: true, message: 'Please input your note!' }] }]"
      />
    </a-form-item>
    <a-form-item label="密码" :label-col="{ span: 5 }" :wrapper-col="{ span: 12 }">
      <a-input
        v-decorator="['password', { rules: [{ required: true, message: 'Please input your note!' }] }]"
      />
    </a-form-item>
    <a-form-item :wrapper-col="{ span: 12, offset: 5 }">
      <a-button type="primary" html-type="submit">登陆</a-button>
      <a-button type @click="reg">注册</a-button>
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
          this.api.login(values).then(res => {
            this.$router.go(-1);
          });
        }
      });
    },
    handleSelectChange(value) {
      console.log(value);
      this.form.setFieldsValue({
        note: `Hi, ${value === "male" ? "man" : "lady"}!`
      });
    },
    reg() {
      this.$router.push("/login/reg");
    }
  }
};
</script>