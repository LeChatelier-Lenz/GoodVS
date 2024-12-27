import React from 'react';
import MySvg from '../assets/pixeltrue-giveaway.svg';

const MyIcon: React.FC = () => {
    return (
        <div>
            <img
                src={MySvg}
                style={{
                    width:"550px",
                    height:"450px",
                    minHeight: "300px",
                    minWidth: "200px",
                }}
                alt="MyIcon" />
        </div>
    );
};

export default MyIcon;