export default class Node {
    constructor(id, type, data) {
        this.id = id;
        this.children = [];
        this.parent = null;
        this.type = type;
        this.data = data;
    }
    setId(id){
        this.id = id;
    }
    setChildren(children){
        this.children = children;
    }
    setParent(parent){
        this.parent = parent;
    }
    setType(type){
        this.type = type;
    }
    setData(data){
        this.data = data;
    }

    getId(){
        return this.id;
    }
    getParent(){
        return this.parent;
    }
    getChildren(){
        return this.children;
    }
    getType(){
        return this.type;
    }
    getData(){
        return this.data;
    }
    getRoot(){
        if(this.parent == null){
            return this;
        }
        return this.parent.getRoot();
    }

    getParentJson(){
        if(this.parent){
            return this.parent.toJSON();
        }else{
            return null;
        }
    }

    getChildrenJson(){
        let json_children = [];
        for(let key in this.children){
            json_children.push(this.children[key].toJSON());
        }
        return json_children;
    }

    toJSON() {
        for (let key in this.children){
            console.log(this.children[key].id);
        }
        return {
            id: this.id,
            type: this.type,
            data: this.data,
            parent: this.getParentJson(),
            //children: this.getChildrenJson()
        }
    }

}