import React from 'react';

interface CardProps {
  name: string;
  price: string;
  description: string;
  duration: string;
  aosDelay: string;
  onPurchase: () => void;
}

const Card: React.FC<CardProps> = ({ name, price, description, duration, aosDelay, onPurchase }) => {
  return (
    <div className="bg-gray-800 p-6 rounded-lg shadow-lg" data-aos="fade-up" data-aos-delay={aosDelay}>
      <h2 className="text-2xl font-bold mb-2">{name}</h2>
      <p className="text-lg mb-2">{price}</p>
      <p className="text-sm mb-2">{description}</p>
      <p className="text-sm mb-4">{duration}</p>
      <button
        onClick={onPurchase}
        className="bg-blue-500 text-white py-2 px-4 rounded"
      >
        Purchase
      </button>
    </div>
  );
};

export default Card;
