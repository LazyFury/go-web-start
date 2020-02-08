<template>
  <a-form :form="form" @submit="handleSubmit">
    <a-form-item label="昵称" :label-col="{ span: 2 }" :wrapper-col="{ span: 12 }">
      <a-input
        v-decorator="['username', { rules: [{ required: true, message: 'Please input your note!' }] }]"
      />
    </a-form-item>

    <a-form-item label="邮箱" :label-col="{ span: 2 }" :wrapper-col="{ span: 12 }">
      <a-input
        v-decorator="['email', { rules: [{ required: false, message: 'Please input your note!' }] }]"
      />
    </a-form-item>

    <a-form-item label="密码" :label-col="{ span: 2 }" :wrapper-col="{ span: 12 }">
      <a-input
        v-decorator="['password', { rules: [{ required: true, message: 'Please input your note!' }] }]"
      />
    </a-form-item>

    <a-form-item label="状态" :label-col="{ span: 2 }" :wrapper-col="{ span: 12 }">
      <a-switch defaultChecked/>
    </a-form-item>

    <a-form-item :wrapper-col="{ span: 12, offset: 2 }">
      <a-button type="primary" html-type="submit">
        Submit
      </a-button>
    </a-form-item>
  </a-form>
</template>

<script>
export default {
  data() {
    return {
      formLayout: 'horizontal',
      form: this.$form.createForm(this, { name: 'coordinated' }),
      id:""
    };
  },
  created(){
      let {id=""} = this.$util.GetParam()
      this.id = id
      if(!id){
          this.$message.error("id不可空")
          window.history.go(-1)
      }
  },
  methods: {
    handleSubmit(e) {
      e.preventDefault();
      this.form.validateFields((err, values) => {
        if (!err) {
          console.log('Received values of form: ', values);
        }
      });
    },
    handleSelectChange(value) {
      console.log(value);
      this.form.setFieldsValue({
        note: `Hi, ${value === 'male' ? 'man' : 'lady'}!`,
      });
    },
  },
};
</script>