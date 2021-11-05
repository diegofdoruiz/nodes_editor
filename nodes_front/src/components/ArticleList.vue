<template>
    <div>
        <h1>Article List</h1>
        <table class="table table-bordered">
            <thead>
                <tr>
                    <th>Uid</th>
                    <th>Name</th>
                </tr>
            </thead>
            <tbody>
                <tr v-for='article in articleItems' :key='article.id'>
                    <td>{{article.uid}}</td>
                    <td>{{article.name}}</td>
                </tr>
            </tbody>
        </table>
    </div>
</template>

<script>
    import { mapGetters } from 'vuex';
    export default {
        name: 'ArticleList',
        computed: {
            ...mapGetters([
                "articleItems",
            ])
        },
        created() {
            this.$store.dispatch('getArticleItems');
        }
    }
</script>