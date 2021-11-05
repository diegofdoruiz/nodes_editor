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
    //import Node from '../../public/assets/js/node-class.js'
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
                        this.modules_info = res.data;
                        let me = this;

                        let first_module_uid = "";

                        //Format all modules from backend
                        res.data.map(function(module, key) {
                            
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
                                    nodes_json[node.id] = me.formatNode(node);
                                });
                            }

                            nodes_data["data"] = nodes_json;
                            data_modules[module.uid] = nodes_data;
                            dataToImport["drawflow"] = data_modules;
                        });

                        //Add default
                        data_modules["Home"] = {"data":{}};
                        if(Object.keys(dataToImport).length !== 0){
                            //console.log(dataToImport);
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
            formatNode(node){
                var new_node = {};
                new_node["id"]      =   node.id;
                new_node["name"]    =   node.name; 
                new_node["pos_x"]   =   node.pos_x;
                new_node["pos_y"]   =   node.pos_y;
                new_node["class"]   =   node.class;
                new_node["data"]    =   node.data;
                new_node["html"]    =   node.html;
                new_node["typenode"] =  node.typenode;

                var inputs = {};
                var outputs = {};
                node.inputs_outputs.map(function(input_output) {
                    var node_name = input_output.name;
                    var node_type = input_output.type;
                    var node_connections = input_output.connections;
                    var connections = [];
                    if(node_connections){
                        node_connections.map(function(connection) {
                            if( node_type == "input"){
                                connections.push({"node": connection.node_number,"input":connection.port});
                            }else{
                                connections.push({"node": connection.node_number,"output":connection.port});
                            }
                        });
                    }

                    if(node_type == "input"){
                        inputs[node_name] = {"connections": connections};
                    }else{
                        outputs[node_name] = {"connections": connections};
                    }

                });
                new_node["inputs"] = inputs;
                new_node["outputs"] = outputs;
                return new_node;
            },
            createNode(module, node){
                //if(Object.keys(node.data).length !== 0){
                    node["data"]["dgraph.type"] = "Data"
                //}
                let inputs_outputs = [];
                var input_output = {};
                var connections = [];
                var connection = {};
                var i = 0;
                //Format inputs
                for (let key in node.inputs) {
                    input_output = {};
                    input_output["node_id"] = node.id; 
                    input_output["name"] = key;
                    input_output["type"] = "input";
                    connections = [];
                    for(i=0; i<node.inputs[key].connections.length; i++){
                        connection = {};
                        connection["node_number"] = node.inputs[key].connections[i].node;
                        connection["port"] = node.inputs[key].connections[i].input;
                        connection["dgraph.type"] = "Connection";
                        connections.push(connection);
                    }
                    input_output["connections"] = connections;
                    input_output["dgraph.type"] = "InputOutput";
                    inputs_outputs.push(input_output);
                }

                //Format Outputs
                for (let key in node.outputs) {
                    input_output = {};
                    input_output["name"] = key;
                    input_output["type"] = "output";
                    connections = [];
                    for(i=0; i<node.outputs[key].connections.length; i++){
                        connection = {};
                        connection["node_number"] = node.outputs[key].connections[i].node;
                        connection["port"] = node.outputs[key].connections[i].output;
                        connection["dgraph.type"] = "Connection";
                        connections.push(connection);
                    }
                    input_output["connections"] = connections;
                    input_output["dgraph.type"] = "InputOutput";
                    inputs_outputs.push(input_output);
                }

                //Add formated inputs/outputs to node - dgraph.type - module_uid
                node["inputs_outputs"] = inputs_outputs;
                node["dgraph.type"] = "Node";
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
                            if(this.modules_info[current_module_key].nodes){
                                this.modules_info[current_module_key].nodes.push(res.data.node);
                            }else{
                                this.modules_info[current_module_key].nodes = [res.data.node];
                            }
                        }
                    }
                });
            },
            deleteNode(id, module){
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

                if(current_node_key >= 0){
                    let node = this.modules_info[current_module_key].nodes[current_node_key];
                    this.$store.dispatch('deleteNode', node).then(res=>{
                        if(res.status == 202){
                            if(res.data.deleted){
                                this.modules_info[current_module_key].nodes.splice(current_node_key,1);
                            }
                        }
                    });
                }
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


                let output_inputs_outputs_key;
                let input_inputs_outputs_key;

                output_node.inputs_outputs.map(function(output, key3) {
                    if(output.name == connection.output_class){
                        output_inputs_outputs_key = key3;
                    }
                });

                input_node.inputs_outputs.map(function(input, key4) {
                    if(input.name == connection.input_class){
                        input_inputs_outputs_key = key4;
                    }
                });

                let output_connection = {};
                output_connection["node_number"] = connection.input_id;
                output_connection["port"] = connection.input_class;
                output_connection["dgraph.type"] = "Connection";

                let input_connection = {};
                input_connection["node_number"] = connection.output_id;
                input_connection["port"] = connection.output_class;
                input_connection["dgraph.type"] = "Connection";

                let payload = {
                    'output_connection': output_connection,
                    'input_connection': input_connection,
                    'output_input_output_uid': this.modules_info[current_module_key].nodes[output_node_key].inputs_outputs[output_inputs_outputs_key].uid,
                    'input_input_output_uid': this.modules_info[current_module_key].nodes[input_node_key].inputs_outputs[input_inputs_outputs_key].uid
                }

                this.$store.dispatch('createNodeConnections', payload).then(res=>{
                    if(res.status == 201){
                        output_connection["uid"] = res.data.connection_output_uid;
                        input_connection["uid"] = res.data.connection_input_uid;
                        if(this.modules_info[current_module_key].nodes[output_node_key].inputs_outputs[output_inputs_outputs_key].connections){
                            this.modules_info[current_module_key].nodes[output_node_key].inputs_outputs[output_inputs_outputs_key].connections.push(
                                output_connection                                
                            );
                            
                        }else{
                            this.modules_info[current_module_key].nodes[output_node_key].inputs_outputs[output_inputs_outputs_key]["connections"] = [output_connection];
                        }

                        if(this.modules_info[current_module_key].nodes[input_node_key].inputs_outputs[input_inputs_outputs_key].connections){
                            this.modules_info[current_module_key].nodes[input_node_key].inputs_outputs[input_inputs_outputs_key].connections.push(
                                input_connection
                            );
                        }else{
                            this.modules_info[current_module_key].nodes[input_node_key].inputs_outputs[input_inputs_outputs_key]["connections"] = [input_connection];
                        }
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


                let output_input_output_key;
                let input_input_output_key;

                output_node.inputs_outputs.map(function(output, key3) {
                    if(output.name == connection.output_class){
                        output_input_output_key = key3;
                    }
                });

                input_node.inputs_outputs.map(function(input, key4) {
                    if(input.name == connection.input_class){
                        input_input_output_key = key4;
                    }
                });
                
                let output_connections = this.modules_info[current_module_key].nodes[output_node_key].inputs_outputs[output_input_output_key].connections;
                let input_connections = this.modules_info[current_module_key].nodes[input_node_key].inputs_outputs[input_input_output_key].connections;

                let output_connection_key;
                let input_connection_key;

                let output_connection_parent_uid = this.modules_info[current_module_key].nodes[output_node_key].inputs_outputs[output_input_output_key].uid;
                let input_connection_parent_uid = this.modules_info[current_module_key].nodes[input_node_key].inputs_outputs[input_input_output_key].uid;
                
                output_connections.map(function(output_connection, key5) {
                    if(output_connection.node_number == connection.input_id && output_connection.port == connection.input_class){
                        output_connection_key = key5;
                    }
                });

                input_connections.map(function(input_connection, key6) {
                    if(input_connection.node_number == connection.output_id && input_connection.port == connection.output_class){
                        input_connection_key = key6;
                    }
                });
                let payload = {
                    'output_connection': output_connections[output_connection_key], 
                    'input_connection': input_connections[input_connection_key],
                    'output_connection_parent_uid': output_connection_parent_uid,
                    'input_connection_parent_uid': input_connection_parent_uid
                };
                this.$store.dispatch('deleteConnections', payload).then(res=>{
                    if(res.status == 200){
                        this.modules_info[current_module_key].nodes[output_node_key].inputs_outputs[output_input_output_key].connections.splice(output_connection_key,1);
                        this.modules_info[current_module_key].nodes[input_node_key].inputs_outputs[input_input_output_key].connections.splice(input_connection_key,1);
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

                //let variables = {};
                //let constants = {};

                /*for( let key in nodes ){
                    let node = nodes[key];

                    //Format variables
                    if(node.name == "variable"){
                        if(this.variableDataValidate(node.data)){
                            let data = {};
                            data["name"]= node.data.name;
                            if(node.data.value){
                                data["value"] = parseInt(node.data.value);
                            }else{
                                data["value"] = parseInt(0);
                            }
                            variables[node.id] = data;
                        }else{
                            return false;
                        }
                        this.variable_names.push(node.data.name);
                    }

                    //Format constants
                    if(node.name == "number"){
                        if(node.data.value){
                            constants[node.id] = parseInt(node.data.value);
                        }else{
                            constants[node.id] = 0;
                        }
                    }

                }*/

                //console.log(nodes);
                let root = this.getRootNode(nodes);
                let object_root = this.recursiveChildren(root);
                //console.log(object_root); //AQUIIIIII

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

                //console.log(variables);
                //console.log(constants);
            }, 
            variableDataValidate(data){
                if(!data.name){
                    Swal.fire({
                        title: "All variables required a name",
                        showCancelButton: false,
                    });
                    return false;
                }
                if(this.variable_names.indexOf(data.name) > -1){
                    Swal.fire({
                        title: "You can not redefine a variable",
                        showCancelButton: false,
                    });
                    return false;
                }

                return true;
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
                            console.log(body_child);
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
            }
        }
    }
</script>

<style lang="scss" scoped>

</style>