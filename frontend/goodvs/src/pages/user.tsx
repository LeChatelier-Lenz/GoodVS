// import React from 'react';
import {
    Box,
    Typography,
    Card,
    CardContent,
    // List,
    // ListItem,
    // ListItemText,
    Avatar,
    Container,
    Grid2, Divider
} from '@mui/material';
import { ArrowForward as ArrowForwardIcon } from '@mui/icons-material';
// import IconButton from "@mui/material/IconButton";
import Button from "@mui/material/Button";
import {useEffect, useState} from "react";
import { GetFollowingProducts} from "../actions/axios.ts";
import ProductCard from "../components/productItem.tsx";
// import {PostLogout} from "../actions/axios.ts";

export default function User () {
    // 示例数据
    interface product{
        "id":string,
        "name":string,
        "url":string,
        "img_url":string,
        "price":number,
        "title":string,
        "category":string,
        "platform":string
    }

    const user = {
        username: localStorage.getItem('name') || 'John Doe',
        email: localStorage.getItem('email') || 'goodvs@example.com',
    };

    const [followedProducts,setFollowProducts] = useState<product[]>([]);

    useEffect(() => {
        GetFollowingProducts(Number(localStorage.getItem('userID')))
            .then((res) => {
                console.log(res);
                setFollowProducts(res.data);
            }).catch((err) => {
            console.log(err);
        });
    }, []);

    const handleLogout = () => {
        localStorage.removeItem('userID');
        localStorage.removeItem('name');
        localStorage.removeItem('email');
        window.location.href = '/';
    }

    return (
        <Container
            sx = {{
                display: "flex",
                flexDirection: "row",
                // alignItems: "center",
                alignItems:"flex-start",
                verticalAlign: "middle",
                // justifyContent: "center",
                // height: "80%",
                maxWidth: "100%",
                minWidth: "700px",
                minHeight: "800px",
                maxHeight: "1000px",
                gap: { xs: 4, sm: 8 },
                py: { xs: 8, sm: 20 },
                mb: 10,
                // border:"solid"
            }}>
            <Box display="flex" sx={{ height: '100vh',width:"100%" }}>
                {/* 左侧：用户信息框，调整为更宽并可以拉伸 */}
                <Box sx={{ width: '25%', padding: 2, borderRight: 1, borderColor: 'grey.300', display: 'flex', flexDirection: 'column' }}>
                    <Card variant="outlined" sx = {{maxWidth: "100%",display :"flex"}}>
                        <CardContent sx={{ display: 'flex', flexDirection: 'column', gap: 2 ,alignItems:"flex-start" ,maxWidth: "100%",
                            overflow: 'hidden',
                        }}>
                            <Box sx = {{display:"flex", alignItems:"center",  minWidth:"20px", maxWidth:"100%",
                                overflow: 'hidden',
                            }} >
                                <Avatar sx={{ width: 56, height: 56, marginRight: 2 }} />
                                <Box sx={{maxWidth: "100%",display:"flex",flexDirection:"column",
                                    overflow: 'hidden',
                                    textOverflow: 'ellipsis',
                                }}>
                                    <Typography variant="h6"
                                                sx={{
                                                    whiteSpace: 'nowrap',
                                                    display: 'block',
                                                    maxWidth:"100%",
                                                    minWidth:"20px",
                                                    overflow: 'hidden',
                                                    textOverflow: 'ellipsis',
                                                }}
                                    >{user.username}</Typography>
                                    <Typography variant="body2" color="textSecondary"
                                                sx={{
                                                    display: 'block',
                                                    maxWidth:"100%",
                                                    minWidth:"20px",
                                                    overflow: 'hidden',
                                                    textOverflow: 'ellipsis',
                                                }}
                                    >{user.email}</Typography>
                                </Box>
                            </Box>
                            <Divider sx={{ marginTop: 2 }} />
                            <Button
                                // sx={{ marginTop: 2 ,display: 'flex',flexDirection:"row",alignItems:"center", justifyContent:"space-between"}}
                            >
                                <Typography variant="body2">用户信息设置</Typography>
                                {/*<IconButton size="small" sx={{ marginLeft: 1 }}>*/}
                                    <ArrowForwardIcon fontSize="small" />
                                {/*</IconButton>*/}
                            </Button>
                            <Button
                                onClick={handleLogout}
                                // sx={{ marginTop: 2 ,display: 'flex',flexDirection:"row",alignItems:"center", justifyContent:"space-between"}}
                            >
                                <Typography variant="body2">退出登录</Typography>
                                {/*<IconButton size="small" sx={{ marginLeft: 1 }}>*/}
                                    <ArrowForwardIcon fontSize="small" />
                                {/*</IconButton>*/}
                            </Button>
                        </CardContent>
                    </Card>
                </Box>

                {/* 右侧：关注商品列表，包含商品图片 */}
                <Box sx={{ width: '60%', padding: 2}}>
                    <Typography variant="h6" gutterBottom>关注的商品</Typography>
                    <Grid2 container spacing={2}>
                        {followedProducts.map((product, index) => (
                            <Grid2 key={index} size={6}
                            >
                                <ProductCard
                                    id={product.id}
                                    title={product.title}
                                    category={product.category}
                                    url={product.url}
                                    imageUrl={product.img_url}
                                    price={product.price}
                                    platform={product.platform}
                                    productName={product.name}
                                    favorited={true}
                                />
                            </Grid2>
                        ))}
                    </Grid2>
                </Box>
            </Box>
    </Container>
    );
};
