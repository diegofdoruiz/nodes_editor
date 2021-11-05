<template>
  <div id="app">
    <router-view/>
  </div>
</template>
<script>
  export default {
    name: 'App',
    components: {
    },
    mounted() {
      this.loadTemaplateScripts();
    },
    methods: {
      async loadPerfectScrollbar(){
        const ps = document.createElement("script");
        ps.setAttribute(
          "src",
          "assets/js/plugins/perfect-scrollbar.min.js"
        );
        document.body.appendChild(ps);
      },
      async loadSmoothScrollbar(){
        await this.loadPerfectScrollbar()
        const ss = document.createElement("script");
        ss.setAttribute(
          "src",
          "assets/js/plugins/smooth-scrollbar.min.js"
        );
        document.body.appendChild(ss);
      },
      async loadTemaplateScripts(){
        await this.loadSmoothScrollbar()
        const ts = document.createElement("script");
        ts.setAttribute(
          "src",
          "assets/js/soft-ui-dashboard.min.js"
        );
        document.body.appendChild(ts);
      }
    },
  }
</script>

<style>
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  /*text-align: center;*/
  color: #2c3e50;
  /*margin-top: 60px;*/
}
</style>
