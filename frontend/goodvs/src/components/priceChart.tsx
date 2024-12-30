import { useEffect, useState } from "react";
import ReactECharts from "echarts-for-react";
import { Box, Typography } from "@mui/material";

export default function PriceChart({ data }) {
    const [option, setOption] = useState({});

    useEffect(() => {
        if (data && data.length > 0) {
            // 处理数据，将价格和时间分别作为 x 轴和 y 轴的数据
            const prices = data.map((item) => item.price);
            const times = data.map((item) => new Date(item.created_at).toLocaleString());

            // 设置 ECharts 配置项
            setOption({
                title: {
                    text: "Price vs Time",
                    left: "center",
                },
                tooltip: {
                    trigger: "axis",
                },
                xAxis: {
                    type: "category",
                    data: times,
                    axisLabel: {
                        rotate: 45, // 让时间标签倾斜，避免重叠
                    },
                    name: "时间",
                },
                yAxis: {
                    type: "value",
                    name: "价格/元",
                    // unit: "￥",
                },
                series: [
                    {
                        data: prices,
                        type: "line",
                        smooth: true, // 平滑折线
                        lineStyle: {
                            color: "#42A5F5",
                        },
                        itemStyle: {
                            color: "#42A5F5",
                        },
                    },
                ],
            });
        }
    }, [data]);

    return (
        <Box sx={{ width: "800px", padding: 3 }}>
            <ReactECharts option={option} style={{ height: "400px" }} />
        </Box>
    );
};
