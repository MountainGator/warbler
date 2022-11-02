import axios from "axios";

export type UserDetails = {
    id: string;
    username: string;
    email: string;
}

export type Warble = {
    id: string;
    user_id: string;
    content: string;
}

export class ApiService {
    private baseUserURL: string = "http://localhost:5000/user/";

    public async checkLogin() {
        const res = await axios.get("http://localhost:5000/check")
        return res.data
    }

    public async getAllWarbles() {
        const res = await axios.get("http://localhost:5000/all-warbles");
        return res.data
    }

    public async getUserDetails(username: string) {
        const res = await axios.get(this.baseUserURL + "details/" + username);
        return res.data
    }


    public async logout() {
        const res = await axios.get(this.baseUserURL + "logout");
        return res.data
    }

    public async updateUser(update: UserDetails) {
        const res = await axios.patch(this.baseUserURL + "update", update);
        return res.data
    }
    
    public async deleteUser(name: string) {
        const res = await axios.delete(this.baseUserURL + "delete/" + name);
        return res.data
    }

    public async newWarble(warble: Warble) {
        const res = await axios.post(this.baseUserURL + "new-warble", warble);
        return res.data
    }

    public async editWarble(warble: Warble) {
        const res = await axios.patch(this.baseUserURL + "edit-warble", warble);
        return res.data
    }

    public async getAllUserWarbles(name: string) {
        const res = await axios.get(this.baseUserURL + "all-warbles/" + name);
        return res.data
    }

    public async deleteWarble(id: string) {
        const res = await axios.delete(this.baseUserURL + "delete-warble/" + id);
        return res.data
    }
}