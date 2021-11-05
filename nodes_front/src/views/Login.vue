<template>
    <div class="login">
        <Login />
        <Footer />
    </div>
</template>

<script>
    import Login from "../components/Login.vue";
    import Footer from "../components/Footer.vue";
    export default {
        name: "login",
        components: {
            Login,
            Footer
        }
    }
</script>

<style lang="scss" scoped>
</style>