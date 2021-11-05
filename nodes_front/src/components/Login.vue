<template>
    <div>
        <div class="container position-sticky z-index-sticky top-0">
            <div class="row">
            <div class="col-12">
                <!-- Navbar -->
                <nav class="navbar navbar-expand-lg blur blur-rounded top-0 z-index-3 shadow position-absolute my-3 py-2 start-0 end-0 mx-4">
                <div class="container-fluid">
                    <router-link class="navbar-brand font-weight-bolder ms-lg-0 ms-3 " to="/">
                        NodesEdit
                    </router-link>
                    <button class="navbar-toggler shadow-none ms-2" type="button" data-bs-toggle="collapse" data-bs-target="#navigation" aria-controls="navigation" aria-expanded="false" aria-label="Toggle navigation">
                        <span class="navbar-toggler-icon mt-2">
                            <span class="navbar-toggler-bar bar1"></span>
                            <span class="navbar-toggler-bar bar2"></span>
                            <span class="navbar-toggler-bar bar3"></span>
                        </span>
                    </button>
                    <div class="collapse navbar-collapse" id="navigation">
                    <ul class="navbar-nav mx-auto">
                        <li class="nav-item">
                            <router-link 
                                class="nav-link me-2" 
                                to="/">
                                <i class="fa fa-home opacity-6 text-dark me-1"></i>
                                <span class="d-sm-inline d-none">Home</span>
                            </router-link>
                        </li>
                        <li class="nav-item">
                            <router-link class="nav-link me-2" to="/login">
                                <i class="fas fa-key opacity-6 text-dark me-1"></i>
                                Sign In
                            </router-link>
                        </li>
                    </ul>
                    
                    </div>
                </div>
                </nav>
                <!-- End Navbar -->
            </div>
            </div>
        </div>
        <main class="main-content  mt-0">
            <section>
            <div class="page-header min-vh-75">
                <div class="container">
                <div class="row">
                    <div class="col-xl-4 col-lg-5 col-md-6 d-flex flex-column mx-auto">
                    <div class="card card-plain mt-8">
                        <div class="card-header pb-0 text-left bg-transparent">
                        <h3 class="font-weight-bolder text-info text-gradient">Welcome back</h3>
                        </div>
                        <div class="card-body">
                        <form role="form" method="POST" @submit.prevent="processing()">
                            <label>Nombre de usuario</label>
                            <div class="mb-3">
                                <input 
                                    type="text" 
                                    v-bind:class="{ 'form-control':true, 'is-invalid':username_error}"
                                    v-model="username"
                                    name="username"
                                    @keyup="resetField"
                                    placeholder="Username" 
                                    aria-label="Username" 
                                    aria-describedby="username-addon">
                                <div class="invalid-feedback">
                                    <span v-for="(error, key) in errors" :key="key">
                                        <template v-if="error.field == 'username'">
                                            {{ error.message }}
                                        </template>
                                    </span>
                                </div>
                            </div>
                            <label>Contraseña</label>
                            <div class="mb-3">
                                <input 
                                    aria-label="Password" 
                                    aria-describedby="password-addon"
                                    type="password" 
                                    v-bind:class="{ 'form-control':true, 'is-invalid':password_error}" 
                                    v-model="password"
                                    name="password"
                                    @keyup="resetField">
                                <div class="invalid-feedback">
                                    <span v-for="(error, key) in errors" :key="key">
                                        <template v-if="error.field == 'password'">
                                            {{ error.message }}
                                        </template>
                                    </span>
                                </div>
                            </div>
                            <!-- <div class="form-check form-switch">
                            <input class="form-check-input" type="checkbox" id="rememberMe" checked="">
                            <label class="form-check-label" for="rememberMe">Remember me</label>
                            </div> -->
                            <div class="text-center">
                            <button type="submit" class="btn bg-gradient-info w-100 mt-4 mb-0">Iniciar</button>
                            </div>
                        </form>
                        </div>
                        <!-- <div class="card-footer text-center pt-0 px-lg-2 px-1">
                        <p class="mb-4 text-sm mx-auto">
                            Don't have an account?
                            <a href="javascript:;" class="text-info text-gradient font-weight-bold">Sign up</a>
                        </p>
                        </div> -->
                    </div>
                    </div>
                    <div class="col-md-6">
                    <div class="oblique position-absolute top-0 h-100 d-md-block d-none me-n8">
                        <div class="oblique-image bg-cover position-absolute fixed-top ms-auto h-100 z-index-0 ms-n6" style="background-image:url('../assets/img/curved-images/curved6.jpg')"></div>
                    </div>
                    </div>
                </div>
                </div>
            </div>
            </section>
        </main>
    </div>
</template>

<script>
    import { mapGetters } from 'vuex';
    export default {
        name: 'LoginComponent',
        data(){
            return {
                username:'',
                password:'',
            }
        },
        computed: {
            ...mapGetters([
                "errors","username_error","password_error"
            ])
        },
        methods:{
            resetField:function(e){
                this.$store.dispatch('ResetField',e.target.name);
            },
            processing(){
                let payload = {'username': this.username, 'password': this.password}
                this.$store.dispatch('signIn',payload);
            }
        }
    }
</script>

<style lang="scss" scoped>

</style>