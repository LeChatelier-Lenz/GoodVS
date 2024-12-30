// ProductCard.tsx
import React from 'react';
import {
    Card,
    CardContent,
    Typography,
    CardMedia,
    Button,
    Tooltip,
    Box,
    DialogTitle,
    DialogContent, DialogContentText,
    DialogActions,
    Dialog
} from '@mui/material';
import LocalOfferIcon from '@mui/icons-material/LocalOffer';
import IconButton from "@mui/material/IconButton";
import StarIcon from '@mui/icons-material/Star';
import Snackbar from "@mui/material/Snackbar";
import ArrowIcon from '@mui/icons-material/ArrowForward';
import {FollowProduct, GetProductPriceList} from "../actions/axios.ts";
import ShowChartIcon from '@mui/icons-material/ShowChart';
// import {PriceChange} from "@mui/icons-material";
import PriceChart from "./priceChart.tsx";

interface ProductCardProps {
    id: string;
    url: string;
    imageUrl: string;
    productName: string;
    price: number;
    platform: string;
    title: string;
    category: string;
    favorited: boolean;
}

const ProductCard: React.FC<ProductCardProps> = ({id, imageUrl, productName,url,price,title,category,favorited}) => {
    const [isFavorite, setIsFavorite] = React.useState(favorited);

    const [click, setClick] = React.useState(false);

    const [open, setOpen] = React.useState(false);

    const [priceList, setPriceList] = React.useState([]);

    const handleClickOpen = () => {
        setOpen(true);
        GetProductPriceList(id).
        then((res) => {
            console.log(res);
            setPriceList(res.data);
        }).catch(
            (err) => {
                console.log(err);
                alert("Failed to get price list");
            }
        );
    };

    const handleClose = () => {
        setOpen(false);
    };

    const handleFavorite = () => {
        const userID = localStorage.getItem('userID');
        if (userID === null) {
            alert("Please sign in first");
            window.location.href = '/signin';
            return;
        }
        if(!isFavorite) {
            // 此时未收藏
            FollowProduct(Number(localStorage.getItem('userID')), id).then(
                (res) => {
                    console.log(res);
                    setIsFavorite(!isFavorite);
                    // setOpen(true);
                    setClick(true);
                }
            ).catch(
                (err) => {
                    console.log(err);
                    alert("Failed to follow product");
                }
            )
        }else{
            // 此时已收藏
            FollowProduct(Number(localStorage.getItem('userID')), id).then(
                (res) => {
                    console.log(res);
                    setIsFavorite(!isFavorite);
                    // setOpen(true);
                    setClick(true);
                }
            ).catch(
                (err) => {
                    console.log(err);
                    alert("Failed to unfollow product");
                }
            )
        }
    }

    return (
        <Card  sx={{ maxWidth: 345, marginY: 1, boxShadow: 3,maxHeight:800,minHeight:400}}>
            <Snackbar
                anchorOrigin={{ vertical:'top',horizontal:'center'}}
                open={click}
                onClose={() => setClick(false)}
                message={isFavorite ? '已收藏' : '已取消收藏'}
                key={'top' + 'center'}
            />
            <CardMedia
                component="img"
                height="180"
                image={imageUrl}
                alt={productName}
                sx={{ objectFit: 'contain',marginTop:2}}
                onClick={
                    () => {
                        window.open(url, '_blank');
                    }
                }

            />
            <CardContent>
                <Box
                    sx={{
                        display: 'flex',
                        flexDirection: 'row',
                        alignItems: 'center',
                        justifyContent: 'space-between',
                        margin: '10px 0'
                    }}
                >
                <Typography variant="body2"  component="div"
                    sx={{
                        borderRadius: '5px',
                        fontSize: 20,
                        fontFamily: 'cursive',
                        backgroundColor: '#dd1111',
                        color: 'white',
                        paddingX: '5px',
                        }}
                >￥{price}</Typography>
                <Typography variant="body1" component="div">
                    {productName}
                </Typography>
                </Box>
                <Box
                    sx={{
                        display: 'flex',
                        flexDirection: 'row',
                        alignItems: 'center',
                        justifyContent: 'flex-start',
                        margin: '10px 0'
                    }}
                >
                    <LocalOfferIcon />
                    <Typography variant="body2" color="text.secondary">
                        {category}
                    </Typography>
                </Box>
                <Tooltip title={title} arrow placement="top">
                <Typography variant="body2" color="text.secondary"
                    sx = {{
                        fontSize: 12,
                        // whiteSpace: "nowrap",
                        // overflowY: "ellipsis",
                        display : "block",
                        overflow: 'hidden',
                        textOverflow: "ellipsis",
                        maxWidth: "100%",
                        maxHeight: "35px",
                        // border: "solid",
                    }}
                >
                    {title}
                </Typography>
                </Tooltip>
                <Box
                    sx={{
                        display: 'flex',
                        flexDirection: 'row',
                        alignItems: 'center',
                        justifyContent: 'space-between',
                        marginTop: '10px'
                    }}
                    >
                    <IconButton onClick={handleFavorite}>
                        <StarIcon color={
                            isFavorite ?  "warning":"disabled"
                        } />
                    </IconButton>
                    <IconButton onClick={handleClickOpen}>
                        <ShowChartIcon/>
                    </IconButton>
                    <Button
                        size="small"
                        color="primary"
                        href={url}
                        target="_blank"
                        rel="noopener noreferrer"
                        sx={{
                            backgroundColor: '#11ccff',
                            color: 'white',
                            fontSize: 12,
                            padding: '5px 10px',
                            borderRadius: '5px',
                            textTransform: 'capitalize',
                        }}
                    >
                        点击查看商品<ArrowIcon/>
                    </Button>
                </Box>
                <Dialog
                    open={open}
                    onClose={handleClose}
                    aria-labelledby="alert-dialog-title"
                    aria-describedby="alert-dialog-description"
                    fullWidth={true}
                    maxWidth={"md"}
                >
                    <DialogTitle id="alert-dialog-title">
                        {productName} 价格变化
                    </DialogTitle>
                    <DialogContent
                        sx={{
                            display: 'flex',
                            flexDirection: 'column',
                            alignItems: 'center',
                            justifyContent: 'center',
                            padding: '10px 20px'
                        }}
                    >
                        <DialogContentText id="alert-dialog-description">
                            <Typography variant="h6" gutterBottom>
                                商品 : {productName} <br/>
                                最新价格 : ￥{price}
                            </Typography>
                        </DialogContentText>
                        <PriceChart data={priceList}/>
                    </DialogContent>
                    <DialogActions>
                        {/*<Button onClick={handleClose}>Disagree</Button>*/}
                        <Button onClick={handleClose} autoFocus>
                            关闭
                        </Button>
                    </DialogActions>
                </Dialog>
            </CardContent>
        </Card>
    );
};

export default ProductCard;
