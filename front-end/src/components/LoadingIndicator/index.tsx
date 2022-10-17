import React from 'react';
import { useSpring, animated } from 'react-spring';
import { percentageToHex } from '../../tokens/Colors';

const LoadingIndicatorShape = ({ isWhite = false }: { isWhite?: boolean }) => {
    const radius = 8;
    const circumfrence = 2 * Math.PI * radius * 8;
    const boxDiameter = '0 0 24 24';
    const boxCenter = 12;
    const none = 'none';
    const foregroundColor = isWhite ? 'white' : '#5B8FDD' + percentageToHex(60);
    const startLength = 1;
    const endLength = 10;

    const width = 3;

    const strokeDashoffsetStyle = useSpring({
        loop: true,
        from: {
            strokeDashoffset: circumfrence - (circumfrence * startLength) / 100,
        },
        to: [
            {
                strokeDashoffset: circumfrence - (circumfrence * endLength) / 100,
            },
            {
                strokeDashoffset: circumfrence - (circumfrence * startLength) / 100,
            },
        ],
        config: {
            duration: 600,
        },
    });

    return (
        <svg viewBox={boxDiameter} fill={none}>
            <animated.circle
                cx={boxCenter}
                cy={boxCenter}
                stroke={foregroundColor}
                strokeWidth={width}
                strokeLinecap={'round'}
                strokeDasharray={circumfrence}
                strokeDashoffset={strokeDashoffsetStyle.strokeDashoffset}
                r={radius}
            />
        </svg>
    );
};

const LoadingIndicator = ({ size, isWhite = false }: { size: string; isWhite?: boolean }) => {
    const rotateStyle = useSpring({
        loop: true,
        from: { transform: 'rotate(0deg)' },
        to: { transform: 'rotate(360deg)' },
        config: {
            duration: 500,
        },
    });

    return (
        <animated.div
            style={{
                transform: rotateStyle.transform,
                width: size,
                height: size,
            }}>
            <LoadingIndicatorShape isWhite={isWhite} />
        </animated.div>
    );
};

export default LoadingIndicator;
