<template>
    <div>
        <Aside />
        <main class="main-content position-relative max-height-vh-100 h-100 mt-1 border-radius-lg ">
            <NavBar />
            <div class="container-fluid py-4">
                <div class="row">
                    <div class="col-12">
                        <Drawflow />
                    </div>
                </div>
                <Footer />
            </div>
        </main>
        <Sidenav />
    </div>
</template>

<script>
    import Aside from '@/components/Aside.vue'
    import NavBar from '@/components/NavBar.vue'
    import Drawflow from '@/components/Drawflow.vue'
    import Footer from '@/components/Footer.vue'
    import Sidenav from '@/components/Sidenav.vue'
    export default {
        components: {
            Aside,
            NavBar,
            Drawflow,
            Footer,
            Sidenav
        }
    }
</script>

<style lang="scss" scoped>

</style>