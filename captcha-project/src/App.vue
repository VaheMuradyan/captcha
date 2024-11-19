<template>
  <div id="app">
    <h1>CAPTCHA Test</h1>
    <div>
      <img :src="captchaImage" alt="CAPTCHA" />
    </div>
    <div>
      <input v-model="userInput" placeholder="Enter CAPTCHA" />
      <button @click="verifyCaptcha">Verify</button>
    </div>
    <button @click="loadCaptcha">Refresh CAPTCHA</button>
  </div>
</template>

<script>
import axios from "axios";

export default {
  data() {
    return {
      captchaID: "",
      captchaImage: "",
      userInput: "",
    };
  },
  methods: {
    async loadCaptcha() {
      try {
        // Ստանում ենք CAPTCHA ID
        const response = await axios.get("http://localhost:8080/captcha");
        this.captchaID = response.data.id;

        // Ստեղծում ենք պատկերի URL
        this.captchaImage = `http://localhost:8080/captcha/image/${this.captchaID}`;
      } catch (error) {
        console.error("Error loading CAPTCHA:", error);
      }
    },
    async verifyCaptcha() {
      try {
        // Ստուգում ենք օգտատիրոջ պատասխանը
        const response = await axios.post("http://localhost:8080/captcha/verify", {
          id: this.captchaID,
          answer: this.userInput,
        });

        alert(response.data.message);
      } catch (error) {
        alert("CAPTCHA incorrect or something went wrong.");
        console.error("Error verifying CAPTCHA:", error);
      }
    },
  },
  mounted() {
    this.loadCaptcha(); // Էջի բեռնումից հետո գեներացնում ենք CAPTCHA
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
