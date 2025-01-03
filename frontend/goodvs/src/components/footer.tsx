import {Container, Link} from "@mui/material";
import React from "react";
import {Divider} from "@mui/material";
import Box from "@mui/material/Box";
import Typography from "@mui/material/Typography";
import InputLabel from "@mui/material/InputLabel";
import Stack from "@mui/material/Stack";
import TextField from "@mui/material/TextField";
import Button from "@mui/material/Button";


export default function Footer() {
    return (
        <React.Fragment>
            <Divider />
            <Container
                sx={{
                    display: 'flex',
                    flexDirection: 'column',
                    alignItems: 'center',
                    gap: { xs: 4, sm: 8 },
                    py: { xs: 8, sm: 10 },
                    textAlign: { sm: 'center', md: 'left' },
                }}
            >
                <Box
                    sx={{
                        display: 'flex',
                        flexDirection: { xs: 'column', sm: 'row' },
                        width: '100%',
                        justifyContent: 'space-between',
                    }}
                >
                    <Box
                        sx={{
                            display: { xs: 'none', sm: 'flex' },
                            flexDirection: 'column',
                            gap: 4,
                            minWidth: { xs: '100%', sm: '60%' },
                        }}
                    >
                        <Box sx={{ width: { xs: '100%', sm: '60%' } }}>
                            <Typography
                                variant="body2"
                                gutterBottom
                                sx={{ fontWeight: 600, mt: 2 }}
                            >
                                订阅 GoodVS
                            </Typography>
                            <Typography variant="body2" sx={{ color: 'text.secondary', mb: 2 }}>
                                订阅新闻通讯，了解我们的最新消息
                            </Typography>
                            <InputLabel htmlFor="email-newsletter">Email | 邮箱</InputLabel>
                            <Stack direction="row" spacing={1} useFlexGap>
                                <TextField
                                    id="email-newsletter"
                                    hiddenLabel
                                    size="small"
                                    variant="outlined"
                                    fullWidth
                                    aria-label="Enter your email address"
                                    placeholder="Your email address"
                                    slotProps={{
                                        htmlInput: {
                                            autoComplete: 'off',
                                            'aria-label': 'Enter your email address',
                                        },
                                    }}
                                    sx={{ width: '250px' }}
                                />
                                <Button
                                    variant="contained"
                                    color="primary"
                                    size="small"
                                    sx={{ flexShrink: 0 }}
                                >
                                    订 阅
                                </Button>
                            </Stack>
                        </Box>
                    </Box>
                    <Box
                        sx={{
                            display: { xs: 'none', sm: 'flex' },
                            flexDirection: 'column',
                            gap: 1,
                        }}
                    >
                        <Typography variant="body2" sx={{ fontWeight: 'medium' }}>
                            Product | 产品
                        </Typography>
                        <Link color="text.secondary" variant="body2" href="#">
                            Features | 敬请期待
                        </Link>
                        <Link color="text.secondary" variant="body2" href="#">
                            Testimonials | 敬请期待
                        </Link>
                        <Link color="text.secondary" variant="body2" href="#">
                            Highlights | 敬请期待
                        </Link>
                        <Link color="text.secondary" variant="body2" href="#">
                            Pricing | 敬请期待
                        </Link>
                        <Link color="text.secondary" variant="body2" href="#">
                            FAQs | 敬请期待
                        </Link>
                    </Box>
                    <Box
                        sx={{
                            display: { xs: 'none', sm: 'flex' },
                            flexDirection: 'column',
                            gap: 1,
                        }}
                    >
                        <Typography variant="body2" sx={{ fontWeight: 'medium' }}>
                            Corperation | 组织
                        </Typography>
                        <Link color="text.secondary" variant="body2" href="/root/about">
                            About us
                        </Link>
                    </Box>
                    <Box
                        sx={{
                            display: { xs: 'none', sm: 'flex' },
                            flexDirection: 'column',
                            gap: 1,
                        }}
                    >
                        <Typography variant="body2" sx={{ fontWeight: 'medium' }}>
                            Legal | 网站政策
                        </Typography>
                        <Link color="text.secondary" variant="body2" href="#">
                            Terms | 条款
                        </Link>
                        <Link color="text.secondary" variant="body2" href="#">
                            Privacy | 隐私
                        </Link>
                        <Link color="text.secondary" variant="body2" href="#">
                            Contact | 联系
                        </Link>
                    </Box>
                </Box>
                {/*<Box*/}
                {/*    sx={{*/}
                {/*        display: 'flex',*/}
                {/*        justifyContent: 'space-between',*/}
                {/*        pt: { xs: 4, sm: 8 },*/}
                {/*        width: '100%',*/}
                {/*        borderTop: '1px solid',*/}
                {/*        borderColor: 'divider',*/}
                {/*    }}*/}
                {/*>*/}
                {/*    <div>*/}
                {/*        <Link color="text.secondary" variant="body2" href="#">*/}
                {/*            Privacy Policy | 隐私政策*/}
                {/*        </Link>*/}
                {/*        <Typography sx={{ display: 'inline', mx: 0.5, opacity: 0.5 }}>*/}
                {/*            &nbsp;•&nbsp;*/}
                {/*        </Typography>*/}
                {/*        <Link color="text.secondary" variant="body2" href="#">*/}
                {/*            Terms of Service | 服务条款*/}
                {/*        </Link>*/}
                {/*    </div>*/}
                    {/*<Stack*/}
                    {/*    direction="row"*/}
                    {/*    spacing={1}*/}
                    {/*    useFlexGap*/}
                    {/*    sx={{ justifyContent: 'left', color: 'text.secondary' }}*/}
                    {/*>*/}
                    {/*    <IconButton*/}
                    {/*        color="inherit"*/}
                    {/*        size="small"*/}
                    {/*        href="https://github.com/LeChatelier-Lenz/hand-chainrity-demo"*/}
                    {/*        aria-label="GitHub"*/}
                    {/*        sx={{ alignSelf: 'center' }}*/}
                    {/*    >*/}
                    {/*    </IconButton>*/}
                    {/*    <IconButton*/}
                    {/*        color="inherit"*/}
                    {/*        size="small"*/}
                    {/*        href="https://x.com/MaterialUI"*/}
                    {/*        aria-label="X"*/}
                    {/*        sx={{ alignSelf: 'center' }}*/}
                    {/*    >*/}
                    {/*    </IconButton>*/}
                    {/*    <IconButton*/}
                    {/*        color="inherit"*/}
                    {/*        size="small"*/}
                    {/*        href="https://www.linkedin.com/company/mui/"*/}
                    {/*        aria-label="LinkedIn"*/}
                    {/*        sx={{ alignSelf: 'center' }}*/}
                    {/*    >*/}
                    {/*    </IconButton>*/}
                    {/*</Stack>*/}
                {/*</Box>*/}
            </Container>
        </React.Fragment>
    );
}