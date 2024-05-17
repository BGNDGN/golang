<template>
  <div>
    <button class="increment-btn" title="buton1" @click="arttirma">+</button> 
    <button class="decrement-btn" title="buton2" @click="azaltma">-</button> 
    <p>{{ counter }}</p> 
  </div>
</template>

<script>
import axios from 'axios';

export default {
  data() {
    return {
      counter: 0 
    };
  },
  methods: {
    async arttirma() {
      await this.counterarttir();
    },
    async azaltma() {
      await this.counterazalt();
    },

    async counterarttir() {
      axios.get('http://localhost:8080/increment')
        .then(response => {
          this.counter = response.data.counter; 
          console.log(response.data);
        })
        .catch(error => {
          console.error('İstek hatası:', error);
        });
    },

    async counterazalt() {
      axios.get('http://localhost:8080/decrement')
        .then(response => {
          this.counter = response.data.counter; 
          console.log(response.data);
        })
        .catch(error => {
          console.error('İstek hatası:', error);
        });
    }
  }
};
</script>
