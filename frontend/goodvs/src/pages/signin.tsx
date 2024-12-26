
import { Box, Button, Card, FormControl, FormLabel, TextField, Typography } from '@mui/material';
import React from "react";
import {PostSignIn} from "../actions/axios.ts";

export default function SignIn() {
    const [emailError, setEmailError] = React.useState(false);
    const [emailErrorMessage, setEmailErrorMessage] = React.useState('');
    const [passwordError, setPasswordError] = React.useState(false);
    const [passwordErrorMessage, setPasswordErrorMessage] = React.useState('');
    const [nameError, setNameError] = React.useState(false);
    const [nameErrorMessage, setNameErrorMessage] = React.useState('');

    // 验证输入
    const validateInputs = () => {
        const name = document.getElementById('name') as HTMLInputElement;
        const email = document.getElementById('email') as HTMLInputElement;
        const password = document.getElementById('password') as HTMLInputElement;
        console.log(
            name.value,
            email.value,
            password.value,
        );

        let isValid = true;

        if (!name.value || name.value.length < 1) {
            setNameError(true);
            setNameErrorMessage('Name is required.');
            isValid = false;
        } else {
            setNameError(false);
            setNameErrorMessage('');
        }

        // 验证邮箱
        if (!email.value || !/\S+@\S+\.\S+/.test(email.value)) {
            setEmailError(true);
            setEmailErrorMessage('请输入正确的邮箱信息');
            isValid = false;
        } else {
            setEmailError(false);
            setEmailErrorMessage('');
        }

        // 验证密码
        if (!password.value || password.value.length < 6) {
            setPasswordError(true);
            setPasswordErrorMessage('Password must be at least 6 characters long.');
            isValid = false;
        } else {
            setPasswordError(false);
            setPasswordErrorMessage('');
        }

        return isValid;
    };


    // 提交表单
    const handleSubmit = async (event: React.FormEvent<HTMLFormElement>) => {
        event.preventDefault();
        const data = new FormData(event.currentTarget);
        let name = data.get('name') as string;
        let email = data.get('email') as string;
        let password = data.get('password') as string;

        // 输入验证
        if (!validateInputs()) {
            return; // 如果验证失败，停止提交
        }

        try{
            PostSignIn(email, password)
            .then((response) => {
                console.log(response);
                if (response.status === 200) {
                    console.log('登录成功');
                    console.log("userID",response.data);
                    alert('登录成功');
                    window.location.href = '/'; // 登录成功，跳转到首页
                }
            }).catch((error) => {
                console.log(error);
                alert('登录失败');
            });
        }catch (e) {
            console.log(e);
            alert('登录失败');
        }

        console.log({
            name: name,
            email: email,
            password: password,
        });
    };


    return (
        <Card variant="outlined"
              sx={{
                  width: '100%',
                  padding: 2, maxWidth: 400, margin: 'auto',
                  align:"center", marginTop: 10 ,alignSelf: 'center',
                  verticalAlign: 'middle'
              }}>
            <Typography
                component="h1"
                variant="h4"
                align="center"
                gutterBottom
                sx={{ fontWeight: 'bold' }}
            >
                用户登录
            </Typography>
            <Box
                component="form"
                onSubmit={handleSubmit}
                sx={{ display: 'flex', flexDirection: 'column', gap: 2 }}
            >
                <FormControl>
                    <FormLabel htmlFor="name">用户名</FormLabel>
                    <TextField
                        autoComplete="name"
                        name="name"
                        required
                        fullWidth
                        id="name"
                        placeholder="用户名"
                        error={nameError}
                        helperText={nameErrorMessage}

                    />
                </FormControl>
                <FormControl>
                    <FormLabel htmlFor="email">邮箱</FormLabel>
                    <TextField
                        required
                        fullWidth
                        id="email"
                        placeholder="your@email.com"
                        name="email"
                        autoComplete="email"
                        variant="outlined"
                        error={emailError}
                        helperText={emailErrorMessage}
                        color={passwordError ? 'error' : 'primary'}
                    />
                </FormControl>
                <FormControl>
                    <FormLabel htmlFor="password">密码</FormLabel>
                    <TextField
                        required
                        fullWidth
                        id="password"
                        placeholder="Password"
                        name="password"
                        type="password"
                        autoComplete="new-password"
                        variant="outlined"
                        error={passwordError}
                        helperText={passwordErrorMessage}
                        color={passwordError ? 'error' : 'primary'}
                    />
                </FormControl>
                <Button
                    type="submit"
                    variant="contained"
                    sx={{ width: '100%', marginTop: 2 }}
                    onClick={validateInputs}
                >
                    登 录
                </Button>
            </Box>
        </Card>
    );
}