import { UserDto } from "@/dtos/users/UserDto";
import { createSlice, PayloadAction } from "@reduxjs/toolkit";




export interface UserState {
    user: UserDto | null
}


const initialState: UserState = {
    user: null
}

const userSlice = createSlice({
    name: 'userState',
    initialState,
    reducers: {
        setUser(state, action: PayloadAction<UserDto>) {
            state.user = action.payload;
        }
    }
})

export const { setUser } = userSlice.actions;
export default userSlice.reducer;