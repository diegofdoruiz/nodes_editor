<template>
    <div class="card p-0">
        <div class="card-body p-0">
            <div class="wrapper">
                <div class="col">
                    <div class="drag-drawflow" draggable="true" @dragstart="mydrawflow.drag($event)" data-node="number">
                        <i class="fas fa-hashtag"></i><span> Number</span>
                    </div>
                    <div class="drag-drawflow" draggable="true" @dragstart="mydrawflow.drag($event)" data-node="variable">
                        <i class="fas fa-subscript"></i><span> Variable</span>
                    </div>
                    <div class="drag-drawflow" draggable="true" @dragstart="mydrawflow.drag($event)" data-node="assign">
                        <i class="fas fa-equals"></i><span> Assign</span>
                    </div>
                    <div class="drag-drawflow" draggable="true" @dragstart="mydrawflow.drag($event)" data-node="addition">
                        <i class="fas fa-plus"></i><span> Addition</span>
                    </div>
                    <div class="drag-drawflow" draggable="true" @dragstart="mydrawflow.drag($event)" data-node="subtraction">
                        <i class="fas fa-minus"></i><span> Subtraction</span>
                    </div>
                    <div class="drag-drawflow" draggable="true" @dragstart="mydrawflow.drag($event)" data-node="multiplication">
                        <i class="fas fa-times"></i><span> Multiplication</span>
                    </div>
                    <div class="drag-drawflow" draggable="true" @dragstart="mydrawflow.drag($event)" data-node="division">
                        <i class="fas fa-divide"></i><span> Division</span>
                    </div>
                    <div class="drag-drawflow" draggable="true" @dragstart="mydrawflow.drag($event)" data-node="ifstatement">
                        <i class="fas fa-project-diagram"></i><span> If statement</span>
                    </div>
                    <div class="drag-drawflow" draggable="true" @dragstart="mydrawflow.drag($event)" data-node="comparation">
                        <i class="fas fa-balance-scale-left"></i><span> Comparation</span>
                    </div>
                    <div class="drag-drawflow" draggable="true" @dragstart="mydrawflow.drag($event)" data-node="myfor">
                        <i class="fas fa-retweet"></i><span> While</span>
                    </div>
                    <!--<div class="drag-drawflow" draggable="true" @dragstart="mydrawflow.drag($event)" data-node="facebook">
                        <i class="fab fa-facebook"></i><span> Facebook</span>
                    </div>
                    <div class="drag-drawflow" draggable="true" @dragstart="mydrawflow.drag($event)" data-node="slack">
                        <i class="fab fa-slack"></i><span> Slack recive message</span>
                    </div>
                    <div class="drag-drawflow" draggable="true" @dragstart="mydrawflow.drag($event)" data-node="github">
                        <i class="fab fa-github"></i><span> Github Star</span>
                    </div>
                    <div class="drag-drawflow" draggable="true" @dragstart="mydrawflow.drag($event)" data-node="telegram">
                        <i class="fab fa-telegram"></i><span> Telegram send message</span>
                    </div>
                    <div class="drag-drawflow" draggable="true" @dragstart="mydrawflow.drag($event)" data-node="aws">
                        <i class="fab fa-aws"></i><span> AWS</span>
                    </div>
                    <div class="drag-drawflow" draggable="true" @dragstart="mydrawflow.drag($event)" data-node="log">
                        <i class="fas fa-file-signature"></i><span> File Log</span>
                    </div>
                    <div class="drag-drawflow" draggable="true" @dragstart="mydrawflow.drag($event)" data-node="google">
                        <i class="fab fa-google-drive"></i><span> Google Drive save</span>
                    </div>
                    <div class="drag-drawflow" draggable="true" @dragstart="mydrawflow.drag($event)" data-node="email">
                        <i class="fas fa-at"></i><span> Email send</span>
                    </div>
                    <div class="drag-drawflow" draggable="true" @dragstart="mydrawflow.drag($event)" data-node="template">
                        <i class="fas fa-code"></i><span> Template</span>
                    </div>
                    <div class="drag-drawflow" draggable="true" @dragstart="mydrawflow.drag($event)" data-node="multiple">
                        <i class="fas fa-code-branch"></i><span> Multiple inputs/outputs</span>
                    </div>
                    <div class="drag-drawflow" draggable="true" @dragstart="mydrawflow.drag($event)" data-node="personalized">
                        <i class="fas fa-fill"></i><span> Personalized</span>
                    </div>
                    <div class="drag-drawflow" draggable="true" @dragstart="mydrawflow.drag($event)" data-node="dbclick">
                        <i class="fas fa-mouse"></i><span> DBClick!</span>
                    </div>-->
                </div>
                <div class="col-right">
                    <div class="menu">
                        <ul>
                            <li 
                                v-for="mod in modules" 
                                :key="mod.id" 
                                @click="mydrawflow.editor.changeModule(mod.id); mydrawflow.changeModule($event);" :class="mod.selected">
                                {{mod.name}} 
                                <i @click="deleteModule($event, mod.id);" class="fas fa-trash-alt delete-module"></i>
                            </li>
                            <li @click="createModule();" class="fs-4" title="Create module" id="create-module">
                                <i class="far fa-plus-square text-info"></i>
                            </li>
                        </ul>
                    </div>
                    <div id="drawflow" @drop="mydrawflow.drop($event)" @dragover="mydrawflow.allowDrop($event)">
                        <!-- <div class="btn-export" @click="showSwal">Export</div> -->
                        <div class="btn-export" @click="pythonCode">Export</div>
                        <div class="btn-clear" @click="clearModule">Clear</div>
                        <div class="btn-lock">
                        <i id="lock" class="fas fa-lock" @click="mydrawflow.editor.editor_mode='fixed'; mydrawflow.changeMode('lock');"></i>
                        <i id="unlock" class="fas fa-lock-open" @click="mydrawflow.editor.editor_mode='edit'; mydrawflow.changeMode('unlock');" style="display:none;"></i>
                        </div>
                        <div class="bar-zoom">
                        <i class="fas fa-search-minus" @click="mydrawflow.editor.zoom_out()"></i>
                        <i class="fas fa-search" @click="mydrawflow.editor.zoom_reset()"></i>
                        <i class="fas fa-search-plus" @click="mydrawflow.editor.zoom_in()"></i>
                        </div>
                    </div>
                </div>
            </div>
            <div class="row">
                <div class="col-12">
                    <button type="button" class="btn btn-primary d-none" data-bs-toggle="modal" data-bs-target="#exampleModal" id="btn-open-code-modal">
                        Launch demo modal
                    </button>
                    <!-- Modal -->
                    <div class="modal fade" id="exampleModal" tabindex="-1" aria-labelledby="exampleModalLabel" aria-hidden="true">
                        <div class="modal-dialog modal-fullscreen">
                            <div class="modal-content">
                                <div class="modal-header">
                                    <h5 class="modal-title" id="exampleModalLabel">Python Code With Brython</h5>
                                    <button type="button" class="btn-close" data-bs-dismiss="modal" aria-label="Close"></button>
                                </div>
                                <div class="modal-body">
                                    <div class="row border h-100">
                                        <div class="col border-rigth">
                                            <div class="row">
                                                <div class="col-12 border-bottom text-end">
                                                    <button type="button" class="btn btn-info m-2" onclick="brython()">
                                                        Run
                                                    </button>
                                                </div>
                                            </div>
                                            <div class="row">
                                                <div class="col-12">
                                                    <pre><code id="show-code"></code></pre>
                                                </div>
                                            </div>
                                        </div>
                                        <div class="col" id="run-code">
                                            <div id="python-result"></div>
                                        </div>
                                    </div>
                                </div>
                                <div class="modal-footer">
                                    <button type="button" class="btn btn-secondary" data-bs-dismiss="modal">Close</button>
                                </div>
                            </div>
                        </div>
                    </div>
                </div>
            </div>
        </div>
    </div>
</template>

<script> 
    import Drawflow from 'drawflow'
    import Swal from 'sweetalert2'
    import MyDrawflow from '../../public/assets/js/mydrawflow.js'
    export default {
        data(){
            return {
                mydrawflow:null,
                modules:[],
                modules_info: [],
                nodes:{},
                variable_names: []
            }
        },
        mounted() {
            var id = document.getElementById("drawflow");
            const editor = new Drawflow(id);
            this.mydrawflow = new MyDrawflow(editor);
            
            //Events
            var l_editor = this.mydrawflow.editor;
            var me = this;
            this.mydrawflow.editor.on('nodeCreated', function(id) {
                var node = l_editor.getNodeFromId(id);
                var module = l_editor.module;
                me.createNode(module, node);
            });
            this.mydrawflow.editor.on('nodeRemoved', function(id) {
                var module = l_editor.module;
                me.deleteNode(id, module);
            });
            this.mydrawflow.editor.on('nodeMoved', function(id) {
                var module = l_editor.module;
                me.nodeUpdatePosition(module, id);
            });
            this.mydrawflow.editor.on('connectionCreated', function(connection) {
                var module = l_editor.module;
                me.createConnection(module,connection);
            });
            this.mydrawflow.editor.on('connectionRemoved', function(connection) {
                var module = l_editor.module;
                me.removeConnection(module,connection);
            });
            this.mydrawflow.editor.on('nodeDataChanged', function(id) {
                var module = l_editor.module;
                me.nodeUpdateData(module, id);
            });

            //load data to import
            this.prepareImportData();
        },
        methods: {
            showSwal(){
                this.mydrawflow.showSwal()
            },
            prepareImportData(){
                let user = JSON.parse(localStorage.getItem('user'));
                this.$store.dispatch('userGetModules', user.username).then(res=>{
                    if(res.status == 200){
                        let dataToImport = {};
                        var data_modules = {};
                        let current_context = this;

                        let back_modules = JSON.parse(JSON.stringify(res.data));
                        
                        this.modules_info = res.data;

                        let first_module_uid = "";

                        //Format all modules from backend
                        back_modules.map(function(module, key) {
                            
                            //Data for nav tabs
                            //current_context.mydrawflow.editor.addModule(module.uid);
                            var selected = '';
                            if(key==0){
                                first_module_uid = module.uid;
                                selected = 'selected';
                            }
                            current_context.modules.push({
                                id: module.uid,
                                name: module.name,
                                selected:selected
                            });
                            
                            //Data for drawflow
                            var nodes_data = {};
                            var nodes_json = {};
                            
                            if(module.nodes){
                                module.nodes.map(function(node) {
                                    if(node.inputs == null){
                                        node["inputs"] = {};
                                    }
                                    if(node.outputs == null){
                                        node["outputs"] = {};
                                    }
                                    for(let key1 in node.inputs){
                                        if(key1 != "uid" && node.inputs[key1].connections == null){
                                            node.inputs[key1].connections = [];
                                        }
                                    }
                                    for(let key2 in node.outputs){
                                        if(key2 != "uid" && node.outputs[key2].connections == null){
                                            node.outputs[key2].connections = [];
                                        }
                                    }
                                    delete node.inputs["uid"];
                                    delete node.outputs["uid"];
                                    nodes_json[node.id] = node;
                                });
                            }
                            nodes_data["data"] = nodes_json;
                            data_modules[module.uid] = nodes_data;
                            dataToImport["drawflow"] = data_modules;
                        });

                        //Add default
                        data_modules["Home"] = {"data":{}};
                        if(Object.keys(dataToImport).length !== 0){
                            this.mydrawflow.editor.import(dataToImport);
                        }
                        if(first_module_uid != ""){
                            this.mydrawflow.editor.changeModule(first_module_uid);
                        }else{
                            this.mydrawflow.editor.changeModule("Home");
                        }
                    }
                });
            },
            createNode(module, node){
                node["data"]["dgraph.type"] = "Data"; //node data DgraphType 
                for(let key in node.inputs){
                    node.inputs[key]["dgraph.type"] = "ConnectionGroup";
                }
                for(let key in node.outputs){
                    node.outputs[key]["dgraph.type"] = "ConnectionGroup";
                }
                node["inputs"]["dgraph.type"] = "Inputs"; //Node inputs DgraphType
                node["outputs"]["dgraph.type"] = "Outputs"; //Node outputs DgraphType
                node["dgraph.type"] = "DrawflowNode"; //node DgraphType
                node["module_uid"] = module;

                this.$store.dispatch('createNode', node).then(res=>{
                    if(res.status == 201){
                        if(res.data.created){
                            let current_module_key = -1;
                            this.modules_info.map(function(module_info, key) {
                                if(module_info.uid == module){
                                    current_module_key = key;
                                }
                            });  
                            let new_node = res.data.node;
                            if(new_node.inputs == null){
                                new_node["inputs"] = {};
                            }
                            if(new_node.outputs == null){
                                new_node["outputs"] = {};
                            }
                            for(let key in new_node.inputs){
                                if(key != "uid" && new_node.inputs[key].connections == null){
                                    new_node.inputs[key].connections = [];
                                }
                            }
                            for(let key in new_node.outputs){
                                if(key != "uid" && new_node.outputs[key].connections == null){
                                    new_node.outputs[key].connections = [];
                                }
                            }
                            if(this.modules_info[current_module_key].nodes){
                                this.modules_info[current_module_key].nodes.push(new_node);
                            }else{
                                this.modules_info[current_module_key].nodes = [new_node];
                            }
                        }
                    }
                });
            },
            deleteNode(id, module){
                let current_module_key = -1;
                let current_node_key;
                for (let key in this.modules_info) {
                    if(this.modules_info[key].uid == module){
                        current_module_key = key;
                        break;
                    }
                }

                for (let key in this.modules_info[current_module_key].nodes) {
                    if(this.modules_info[current_module_key].nodes[key].id == id){
                        current_node_key = key;
                        break;
                    }
                }
                this.$store.dispatch('deleteNode', this.modules_info[current_module_key].nodes[current_node_key]).then(res=>{
                    if(res.status == 202){
                        if(res.data.deleted){
                            this.modules_info[current_module_key].nodes.splice(current_node_key,1);
                        }
                    }
                });
            },
            createConnection(module, connection){
                //{output_id: '1', input_id: '2', output_class: 'output_1', input_class: 'input_1'}
                let current_module_key = -1;
                let output_node_key = -1;
                let input_node_key = -1;
                for (let key in this.modules_info) {
                    if(this.modules_info[key].uid == module){
                        current_module_key = key;
                        break;
                    }
                }

                for (let key2 in this.modules_info[current_module_key].nodes) {
                    if(this.modules_info[current_module_key].nodes[key2].id == connection.output_id){
                        output_node_key = key2;
                    }
                    if(this.modules_info[current_module_key].nodes[key2].id == connection.input_id){
                        input_node_key = key2;
                    }
                }

                let output_node = this.modules_info[current_module_key].nodes[output_node_key];
                let input_node = this.modules_info[current_module_key].nodes[input_node_key];

                let output_group_connections = output_node.outputs[connection.output_class];
                let input_group_connections = input_node.inputs[connection.input_class];

                let output_connection = {};
                output_connection["node"] = connection.input_id;
                output_connection["output"] = connection.input_class;
                output_connection["group"] = {'uid': output_group_connections.uid};
                output_connection["dgraph.type"] = "Connection";

                let input_connection = {};
                input_connection["node"] = connection.output_id;
                input_connection["input"] = connection.output_class;
                input_connection["group"] = {'uid':input_group_connections.uid};
                input_connection["dgraph.type"] = "Connection";

                let payload = {
                    'output_connection': output_connection,
                    'input_connection': input_connection
                }

                this.$store.dispatch('createNodeConnections', payload).then(res=>{
                    if(res.status == 201){
                        output_connection["uid"] = res.data.connection_output_uid;
                        input_connection["uid"] = res.data.connection_input_uid;
                        output_group_connections.connections.push(output_connection);
                        input_group_connections.connections.push(input_connection);
                    }
                });
            },
            removeConnection(module, connection){
                let current_module_key = -1;
                let output_node_key = -1;
                let input_node_key = -1;
                for (let key in this.modules_info) {
                    if(this.modules_info[key].uid == module){
                        current_module_key = key;
                        break;
                    }
                }

                for (let key2 in this.modules_info[current_module_key].nodes) {
                    if(this.modules_info[current_module_key].nodes[key2].id == connection.output_id){
                        output_node_key = key2;
                    }
                    if(this.modules_info[current_module_key].nodes[key2].id == connection.input_id){
                        input_node_key = key2;
                    }
                }

                let output_node = this.modules_info[current_module_key].nodes[output_node_key];
                let input_node = this.modules_info[current_module_key].nodes[input_node_key];

                let output_group_connections = output_node.outputs[connection.output_class];
                let input_group_connections = input_node.inputs[connection.input_class];

                let output_connection = {};
                let input_connection = {};
                let input_connection_key;
                let output_connection_key;

                for (let key in input_group_connections.connections){
                    if(input_group_connections.connections[key].node == connection.output_id){
                        input_connection = input_group_connections.connections[key];
                        input_connection_key = key;
                    }
                }

                for (let key in output_group_connections.connections){
                    if(output_group_connections.connections[key].node == connection.input_id){
                        output_connection = output_group_connections.connections[key];
                        output_connection_key = key;
                    }
                }

                let payload = {
                    'output_connection': output_connection, 
                    'input_connection': input_connection
                };
                this.$store.dispatch('deleteConnections', payload).then(res=>{
                    if(res.status == 200){
                        input_group_connections.connections.splice(input_connection_key,1);
                        output_group_connections.connections.splice(output_connection_key,1);
                    }
                });
            },
            findNodeInDrawflow(module, id){
                let drawflow_module = this.mydrawflow.editor.drawflow.drawflow[module];
                let drawflow_node = drawflow_module.data[id];
                return drawflow_node;
            },
            nodeUpdatePosition(module, id){
                let drawflow_node = this.findNodeInDrawflow(module, id);
                let current_module_key = -1;
                let current_node_key = -1;
                for (let key in this.modules_info) {
                    if(this.modules_info[key].uid == module){
                        current_module_key = key;
                        break;
                    }
                }
                for (let key in this.modules_info[current_module_key].nodes) {
                    if(this.modules_info[current_module_key].nodes[key].id == id){
                        current_node_key = key;
                        break;
                    }
                }
                let payload = {
                    'uid':this.modules_info[current_module_key].nodes[current_node_key].uid,
                    'pos_x':drawflow_node.pos_x,
                    'pos_y':drawflow_node.pos_y
                }
                this.$store.dispatch('nodeUpdatePosition', payload).then(res=>{
                    if(res.status == 201){
                        if(res.data.updated){
                            this.modules_info[current_module_key].nodes[current_node_key].pos_x=drawflow_node.pos_x;
                            this.modules_info[current_module_key].nodes[current_node_key].pos_y=drawflow_node.pos_y;
                        }
                    }
                });
            },
            nodeUpdateData(module, id){
                let drawflow_node = this.findNodeInDrawflow(module, id);
                let current_module_key = -1;
                let current_node_key = -1;
                for (let key in this.modules_info) {
                    if(this.modules_info[key].uid == module){
                        current_module_key = key;
                        break;
                    }
                }
                for (let key1 in this.modules_info[current_module_key].nodes) {
                    if(this.modules_info[current_module_key].nodes[key1].id == id){
                        current_node_key = key1;
                        break;
                    }
                }

                let payload = {};
                payload["uid"] = this.modules_info[current_module_key].nodes[current_node_key].data.uid;
                for(let key3 in drawflow_node.data){
                    payload[key3] = drawflow_node.data[key3];
                }
                this.$store.dispatch('nodeUpdateData', payload).then(res=>{
                    if(res.status == 202){
                        if(res.data.updated){
                            this.modules_info[current_module_key].nodes[current_node_key].data = payload;
                        }
                    }
                });
            },
            createModule(){
                let user = JSON.parse(localStorage.getItem('user'));
                Swal.fire({
                    title: 'Module name',
                    input: 'text',
                    inputAttributes: {
                        autocapitalize: 'off'
                    },
                    showCancelButton: true,
                    confirmButtonText: 'Create',
                    showLoaderOnConfirm: true,
                    preConfirm: (name) => {
                        let module = {
                            'name': name,
                            'owner': user.username
                        }
                        return this.$store.dispatch('searchModule', module).then(res=>{
                            if(res.success){
                                if(res.isset){
                                    Swal.showValidationMessage(
                                        `A module with this name already exists`
                                    )
                                }
                            }else{
                                Swal.showValidationMessage(
                                    `Request failed`
                                )
                            }
                        });
                    },
                    allowOutsideClick: () => !Swal.isLoading()
                }).then((result) => {
                    if (result.isConfirmed) {
                        var module = {
                            'name': result.value,
                            'owner': user.username
                        }
                        return this.$store.dispatch('createModule', module).then(res=>{
                            if(res.status == 201){
                                this.mydrawflow.editor.addModule(res.data.uid);
                                this.mydrawflow.editor.changeModule(res.data.uid);
                                for(var i=0; i<this.modules.length; i++){
                                    this.modules[i].selected='';
                                }
                                this.modules.push({
                                    id: res.data.uid,
                                    name: module.name,
                                    selected:'selected'
                                });
                                module["uid"] = res.data.uid;
                                module["nodes"] = [];
                                this.modules_info.push(module);
                            }
                        });
                    }
                })
            },
            clearModule(){
                let module = this.mydrawflow.editor.module;
                this.$store.dispatch('clearModule', module).then(res=>{
                    if(res.status == 200){
                        //this.mydrawflow.editor.clearModuleSelected();
                        let current_module_key = -1;
                        for (let key in this.modules_info) {
                            if(this.modules_info[key].uid == module){
                                current_module_key = key;
                                break;
                            }
                        }
                        this.modules_info[current_module_key].nodes = [];
                        this.mydrawflow.editor.precanvas.innerHTML = "";
                        this.mydrawflow.editor.drawflow.drawflow[module] = { "data": {} };
                    }
                });
            },
            deleteModule(event, id){
                let user = JSON.parse(localStorage.getItem('user'));
                event.stopPropagation();
                Swal.fire({
                    icon: 'error',
                    title: 'Deleting module. Are you sure?',
                    inputAttributes: {
                        autocapitalize: 'off'
                    },
                    showCancelButton: true,
                    confirmButtonText: 'Delete',
                    showLoaderOnConfirm: true,
                    allowOutsideClick: () => !Swal.isLoading()
                }).then((result) => {
                    if (result.isConfirmed) {
                        var index;
                        let current_module_key = -1;
                        this.modules.map(function(module, key) {
                            if(module.id == id){
                                index = key;
                            }
                        });
                    
                        for (let key in this.modules_info) {
                            if(this.modules_info[key].uid == id){
                                current_module_key = key;
                                break;
                            }
                        }

                        var module = {
                            uid: id,
                            owner: user.username
                        }
                        return this.$store.dispatch('userDeleteModule', module).then(res=>{
                            if(res.status == 202){
                                this.modules.splice(index, 1);
                                this.modules_info.splice(current_module_key, 1);
                                this.mydrawflow.editor.removeModule(res.data.uid);
                            }
                        });
                    }
                })
            },
            pythonCode(){
                let module = this.mydrawflow.editor.module;
                let drawflow_module = this.mydrawflow.editor.drawflow.drawflow[module];
                let nodes = drawflow_module.data;
                this.variable_names = [];
                this.nodes = nodes;

                let root = this.getRootNode(nodes);
                let object_root = this.recursiveChildren(root);

                let errors = this.treeVerification(object_root);

                if(errors.length > 0){

                    Swal.fire({
                        icon: 'error',
                        title: 'Errors were found...',
                        text: ".",
                    });

                    let html_list = document.createElement("ul");
                    for(let key in errors){
                        let html_item = document.createElement("li");
                        html_item.appendChild(document.createTextNode(errors[key].error));
                        html_list.appendChild(html_item);
                    }
                    document.getElementById("swal2-html-container").appendChild(html_list);
                    return false;
                }

                //Unparse python code
                let line = "from browser import document\n\n";
                line += "def display_result(result):\n";
                line += "   text = document.createTextNode('> '+str(result))\n";
                line += "   document.getElementById(\"python-result\").appendChild(text)\n";
                line += "   new_line = document.createElement('br')\n";
                line += "   document.getElementById(\"python-result\").appendChild(new_line)\n\n";
                let str_python_code = line + this.printTree(object_root);
                let arr_python_code = str_python_code.split("\n");

                //Add code to display
                let show_code = document.getElementById("show-code");
                let run_code = document.getElementById("run-code");

                //Add python script :)
                var old_scripts =  document.querySelectorAll("script[type='text/python']");
                old_scripts.forEach(function(old_script){
                    old_script.parentNode.removeChild(old_script);
                });
                let python_element = document.createElement("script");
                python_element.type = "text/python";
                python_element.appendChild(document.createTextNode(line));
    
                let new_str_python = "";
                for(let key5 in arr_python_code){
                    if(arr_python_code[key5].indexOf(' ') == 0){
                        if(arr_python_code[key5] != ""){
                            new_str_python += "&nbsp;&nbsp;"+arr_python_code[key5] + "<br>";
                        }
                    }else{
                        if(arr_python_code[key5] != ""){
                            new_str_python += arr_python_code[key5] + "<br>";
                        }
                    }
                    line = document.createTextNode(arr_python_code[key5]+"\n");
                    python_element.appendChild(line);
                }
                show_code.innerHTML = new_str_python;
                run_code.appendChild(python_element);

                
                //Open modal code
                let python_result = document.getElementById("python-result");
                python_result.innerHTML = "";
                let btn_open_code_modal = document.getElementById("btn-open-code-modal");
                btn_open_code_modal.click();
            },
            getRootNode(nodes){
                for(let key in nodes){
                    let outputs = nodes[key].outputs;
                    let inputs = nodes[key].inputs;
                    let counter = 0;
                    let counter2 = 0;
                    for(let key2 in outputs){
                        let connections = outputs[key2].connections;
                        if(connections){
                            if(connections.length > 0){
                                counter = counter + 1;
                            }
                        }
                    }
                    for(let key3 in inputs){
                        let connections = inputs[key3].connections;
                        if(connections){
                            if(connections.length > 0){
                                counter2 = counter2 + 1;
                            }
                        }
                    }
                    if(counter === 0 && counter2 > 0){
                        return nodes[key];
                    }
                }  
                return null;
            },
            nodeHasChildren(node){
                let counter = 0;
                for(let key in node.inputs){
                    let connections = node.inputs[key].connections;
                    if(connections){
                        if(connections.length > 0){
                            counter = counter + 1;
                        }
                    }
                }
                if(counter > 0){
                    return true;
                }
                return false;
            },
            nodeGetChildren(object_node){
                let node = this.nodes[object_node.id];

                //First add sequence on inputs
                for(let key0 in node.outputs){
                    let outputs_connections = node.outputs[key0].connections;
                    if(outputs_connections){
                        if(outputs_connections.length > 0){
                            for(let key1 in outputs_connections){
                                let output_connection = outputs_connections[key1];
                                object_node[output_connection.node] = output_connection.output;
                            }
                        }
                    }
                }

                let children = []
                for(let key2 in node.inputs){
                    let connections = node.inputs[key2].connections;
                    if(connections){
                        if(connections.length > 0){
                            for(let key3 in connections){
                                let connection = connections[key3];
                                let child = this.nodes[connection.node];
                                let child_object = this.recursiveChildren(child);
                                child_object["parent"] = object_node;
                                children.push(child_object);
                            }
                        }
                    }
                }
                return children;
            },
            recursiveChildren(node){
                let object_node = {};
                object_node["id"] = node.id;
                object_node["type"] = node.name;
                object_node["data"] = node.data;
                object_node["visited"] = false;
                if(this.nodeHasChildren(node)){
                    let children = this.nodeGetChildren(object_node);
                    object_node["children"] = children;
                }
                return object_node;
            },
            printTree(object_node){
                if(object_node.type == "number"){
                    let value = 0;
                    if(object_node.data.value){
                        value = parseInt(object_node.data.value);
                    }
                    return "constant_"+object_node.id+" = "+value;
                }else if(object_node.type == "variable"){
                    let str_return = "";
                    if(object_node.children){
                        if(object_node.children.length == 1){
                            str_return += this.printTree(object_node.children[0])+"\n";
                            str_return += object_node.data.name+" = constant_"+object_node.children[0].id+"\n";
                            return str_return;
                        }else{
                            return object_node.data.name+" = 0\n";
                        }
                    }
                    return str_return;
                }else if(object_node.type == "addition"){
                    return this.buildBinaryStringReturn(object_node, "+");
                }else if(object_node.type == "subtraction"){
                    return this.buildBinaryStringReturn(object_node, "-");
                }else if(object_node.type == "multiplication"){
                    return this.buildBinaryStringReturn(object_node, "*");
                }else if(object_node.type == "division"){
                    return this.buildBinaryStringReturn(object_node, "/");
                }else if(object_node.type == "comparation"){
                    let operator =  object_node.data.operator;
                    return this.buildBinaryStringReturn(object_node, operator);
                }else if(object_node.type == "assign"){
                    let str_return = "";
                    str_return += this.printTree(object_node.children[0])+"\n";
                    if(object_node.children[0].type == "assign"){
                        str_return += object_node.data.name+" = "+object_node.children[0].data.name+"\n";
                    }else if(object_node.children[0].type == "variable"){
                        str_return += object_node.data.name+" = "+object_node.children[0].data.name+"\n";
                    }else{
                        str_return += object_node.data.name+" = constant_"+object_node.children[0].id+"\n";
                    }
                    str_return += "print("+object_node.data.name+")\n";
                    str_return += "display_result("+object_node.data.name+")\n";
                    return str_return;


                }else if(object_node.type == "ifstatement"){
                    let str_return = "";
                    let comparation_node = null;
                    let if_child = null;
                    let else_child = null;

                    //If statement childs
                    for(let j=0; j<object_node.children.length; j++){
                        let child_node = object_node.children[j];
                        if(child_node[object_node.id] == "input_1"){ //Comparation
                            comparation_node = child_node;
                        }else if(child_node[object_node.id] == "input_2"){ //If Body
                            if_child = child_node;
                        }else{ // Else body
                            else_child = child_node;
                        }
                    }

                    //If comparation
                    let comparation = "";
                    if(comparation_node){
                        let children_0 = comparation_node.children[0];
                        let children_1 = comparation_node.children[1];
                        let children_0_name = "constant_"+children_0.id;
                        let children_1_name = "constant_"+children_1.id;
                        if(children_0.type == "variable"){
                            children_0_name = children_0.data.name;
                        }
                        if(children_1.type == "variable"){
                            children_1_name = children_1.data.name;
                        }
                        comparation = children_0_name+" "+comparation_node.data.operator+" "+children_1_name;
                        str_return += this.printTree(comparation_node)+"\n";
                    }

                    //If body
                    str_return += "if "+comparation+":\n";
                    if(if_child){
                        let str_if_body = this.printTree(if_child);
                        let arr_if_body = str_if_body.split("\n");
                        for(let i = 0; i<arr_if_body.length; i++){
                            str_return += "   "+arr_if_body[i]+"\n";
                        }
                    }else{
                        str_return += "   pass\n";
                    }

                    str_return += "else:\n";
                    if(else_child){
                        let str_else_body = this.printTree(else_child);
                        let arr_else_body = str_else_body.split("\n");
                        for(let i = 0; i<arr_else_body.length; i++){
                            str_return += "   "+arr_else_body[i]+"\n";
                        }
                    }else{
                        str_return += "   pass\n";
                    }                        
                    return str_return;  
                }else if(object_node.type == "myfor"){
                    let str_return = "";
                    let comparation_node = null;
                    let body_child = null;
                    //If statement childs
                    for(let j=0; j<object_node.children.length; j++){
                        let child_node = object_node.children[j];
                        if(child_node[object_node.id] == "input_1"){ //Comparation
                            comparation_node = child_node;
                        }else if(child_node[object_node.id] == "input_2"){ //If Body
                            body_child = child_node;
                        }
                    }

                    //If comparation
                    let comparation = "";
                    if(comparation_node){
                        let children_0 = comparation_node.children[0];
                        let children_1 = comparation_node.children[1];
                        let children_0_name = "constant_"+children_0.id;
                        let children_1_name = "constant_"+children_1.id;
                        if(children_0.type == "variable"){
                            children_0_name = children_0.data.name;
                        }
                        if(children_1.type == "variable"){
                            children_1_name = children_1.data.name;
                        }
                        comparation = children_0_name+" "+comparation_node.data.operator+" "+children_1_name;
                        str_return += this.printTree(comparation_node)+"\n";
                    }

                    //If body
                    str_return += "while "+comparation+":\n";
                    if(body_child){
                        let str_while_body = this.printTree(body_child);
                        let arr_while_body = str_while_body.split("\n");
                        for(let i = 0; i<arr_while_body.length; i++){
                            str_return += "   "+arr_while_body[i]+"\n";
                        }
                    }else{
                        str_return += "   pass\n";
                    }

                    return str_return;
                }
            },
            buildBinaryStringReturn(object_node, operator){
                let str_return = "";
                str_return += this.printTree(object_node.children[0])+"\n";
                str_return += this.printTree(object_node.children[1])+"\n";
                let name_child_0 = "constant_"+object_node.children[0].id;
                let name_child_1 = "constant_"+object_node.children[1].id;
                if(object_node.children[0].type == "variable"){
                    name_child_0 = object_node.children[0].data.name;
                }
                if(object_node.children[1].type == "variable"){
                    name_child_1 = object_node.children[1].data.name;
                }
                if(object_node.type != "comparation"){
                    str_return += "constant_"+object_node.id+" = "+name_child_0+" "+operator+" "+name_child_1;
                }
                return str_return;
            },
            treeVerification(node_tree){
                if(node_tree.type == "number"){
                    let errors = [];
                    if(!node_tree.data.value){
                        errors.push({
                            id: node_tree.id,
                            error: "The node number #"+node_tree.id+" has no value"
                        });
                    }
                    return errors;
                }else if(node_tree.type == "variable"){
                    let errors = [];
                    if(!node_tree.data.name){
                        errors.push({
                            id: node_tree.id,
                            error: "The node  variable #"+node_tree.id+" has no name"
                        });
                    }
                    if(node_tree.children.length > 0){
                        for(let i = 0; i<node_tree.children.length; i++){
                            errors = errors.concat(this.treeVerification(node_tree.children[i]));
                        }
                    }
                    return errors;
                }else if(node_tree.type == "addition" || node_tree.type == "subtraction" || node_tree.type == "multiplication" || node_tree.type == "comparation"){
                    let errors = [];
                    if(node_tree.children.length != 2){
                        errors.push({
                            id: node_tree.id,
                            error: "The node "+node_tree.type+" #"+node_tree.id+" requires two numeric values"
                        });
                    }
                    if(node_tree.children.length > 0){
                        for(let i = 0; i<node_tree.children.length; i++){
                            errors = errors.concat(this.treeVerification(node_tree.children[i]));
                        }
                    }
                    return errors;
                }else if(node_tree.type == "division"){
                    let errors = [];
                    if(node_tree.children.length != 2){
                        errors.push({
                            id: node_tree.id,
                            error: "The node "+node_tree.type+" #"+node_tree.id+" requires two numeric values"
                        });
                    }
                    if(node_tree.children.length == 2){
                        if(node_tree.children[1].type == "number" && (node_tree.children[1].data.value == 0 || node_tree.children[1].data.value == "0")){
                            errors.push({
                                id: node_tree.id,
                                error: "The node "+node_tree.type+" #"+node_tree.id+" has a zero in the divisor"
                            });
                        }
                        for(let i = 0; i<node_tree.children.length; i++){
                            errors = errors.concat(this.treeVerification(node_tree.children[i]));
                        }
                    }
                    return errors;
                }else if(node_tree.type == "assign"){
                    let errors = [];
                    if(node_tree.children.length != 1){
                        errors.push({
                            id: node_tree.id,
                            error: "The node "+node_tree.type+" #"+node_tree.id+" requires one numeric value"
                        });
                    }
                    //Group variables names
                    let module_data = this.mydrawflow.editor.drawflow.drawflow[this.mydrawflow.editor.module].data;
                    let variable_names = [];
                    for (let key in module_data){
                        if(module_data[key].name == "variable"){
                            if(module_data[key].data.name){
                                variable_names.push(module_data[key].data.name);
                            }
                        }
                    }
                    if(node_tree.data.name == "" || !node_tree.data.name){
                        errors.push({
                            id: node_tree.id,
                            error: "The node "+node_tree.type+" #"+node_tree.id+" does not have an associated variable"
                        });
                    }else{
                        if(!variable_names.includes(node_tree.data.name)){
                            errors.push({
                                id: node_tree.id,
                                error: "The node "+node_tree.type+" #"+node_tree.id+" the variable with name '"+node_tree.data.name+"' does not exist",
                            });
                        }
                    }
                    if(node_tree.children.length > 0){
                        for(let i = 0; i<node_tree.children.length; i++){
                            errors = errors.concat(this.treeVerification(node_tree.children[i]));
                        }
                    }
                    return errors;
                }/*else if(node_tree.type == "ifstatement"){
                    
                }else if(node_tree.type == "myfor"){
                    
                }*/
            }
        }
    }
</script>

<style lang="scss" scoped>

</style>