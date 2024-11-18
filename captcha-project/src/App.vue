<template>
  <div id="app">
    <h1>CAPTCHA Test</h1>
    <div>
      <img :src="captchaImage" alt="CAPTCHA" />
    </div>
    <div>
      <input v-model="userInput" placeholder="Մուտքագրեք CAPTCHA-ն" />
      <button @click="verifyCaptcha">Ստուգել</button>
    </div>
    <button @click="loadCaptcha">Փոխել CAPTCHA-ն</button>
  </div>
</template>

<script>
import axios from "axios";

export default {
  data() {
    return {
      captchaId: "",
      captchaImage: "",
      userInput: "",
    };
  },
  methods: {
    async loadCaptcha() {
      try {
        const response = await axios.get("http://localhost:8080/captcha");
        this.captchaId = response.data.id;
        this.captchaImage = response.data.captcha;
      } catch (error) {
        console.error("Error loading CAPTCHA:", error);
      }
    },
    async verifyCaptcha() {
      try {
        const response = await axios.post("http://localhost:8080/verify", {
          id: this.captchaId,
          answer: this.userInput,
        });
        alert(response.data.message);
      } catch (error) {
        alert("CAPTCHA սխալ է կամ խնդիր կա: Խնդրում ենք փորձել նորից:");
        console.error("Error verifying CAPTCHA:", error);
      }
    },
  },
  mounted() {
    this.loadCaptcha();
  },
};
</script>

<style>
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  text-align: center;
  margin-top: 50px;
}
img {
  max-width: 100%;
  height: auto;
  margin-bottom: 20px;
}
</style>
