// ProductCard.tsx
import React from 'react';
import { Card, CardContent, Typography, CardMedia, Button } from '@mui/material';

interface ProductCardProps {
    imageUrl: string;
    productName: string;
    platform: string;
    productLink: string;
}

const ProductCard: React.FC<ProductCardProps> = ({ imageUrl, productName, platform, productLink }) => {
    return (
        <Card sx={{ maxWidth: 345, margin: 2, boxShadow: 3 }}>
            <CardMedia
                component="img"
                height="140"
                image={imageUrl}
                alt={productName}
            />
            <CardContent>
                <Typography variant="h6" component="div">
                    {productName}
                </Typography>
                <Typography variant="body2" color="text.secondary">
                    {platform}
                </Typography>
                <Button
                    size="small"
                    color="primary"
                    href={productLink}
                    target="_blank"
                    rel="noopener noreferrer"
                >
                    查看商品
                </Button>
            </CardContent>
        </Card>
    );
};

export default ProductCard;
