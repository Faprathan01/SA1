import { CreditCardInterface } from "../../interfaces/ICreditCard";
import { PaypalInterface } from "../../interfaces/IPaypal";
import { PromptpayInterface } from "../../interfaces/IPromptpay";

// API URL
const apiUrl = "http://localhost:3036";

// Credit Card Functions
async function GetCreditCards() {
  const requestOptions = {
    method: "GET",
    headers: {
      "Content-Type": "application/json",
    },
  };

  try {
    const res = await fetch(`${apiUrl}/creditcards`, requestOptions);
    if (res.status === 200) {
      return res.json();
    } else {
      return false;
    }
  } catch (error) {
    console.error("Error fetching credit cards:", error);
    return false;
  }
}

async function GetCreditCardById(id: number | undefined) {
  const requestOptions = {
    method: "GET",
    headers: {
      "Content-Type": "application/json",
    },
  };

  try {
    const res = await fetch(`${apiUrl}/creditcards/${id}`, requestOptions);
    if (res.status === 200) {
      return res.json();
    } else {
      return false;
    }
  } catch (error) {
    console.error("Error fetching credit card by ID:", error);
    return false;
  }
}

async function CreateCreditCard(data: CreditCardInterface) {
  const requestOptions = {
    method: "POST",
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify(data),
  };

  try {
    const res = await fetch(`${apiUrl}/creditcards`, requestOptions);
    if (res.status === 201) {
      return res.json();
    } else {
      return false;
    }
  } catch (error) {
    console.error("Error creating credit card:", error);
    return false;
  }
}

async function UpdateCreditCard(data: CreditCardInterface) {
  const requestOptions = {
    method: "PATCH",
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify(data),
  };

  try {
    const res = await fetch(`${apiUrl}/creditcards/${data.CreditCardID}`, requestOptions);
    if (res.status === 200) {
      return res.json();
    } else {
      return false;
    }
  } catch (error) {
    console.error("Error updating credit card:", error);
    return false;
  }
}

async function DeleteCreditCardByID(id: number | undefined) {
  const requestOptions = {
    method: "DELETE",
  };

  try {
    const res = await fetch(`${apiUrl}/creditcards/${id}`, requestOptions);
    if (res.status === 200) {
      return true;
    } else {
      return false;
    }
  } catch (error) {
    console.error("Error deleting credit card:", error);
    return false;
  }
}

// PayPal Functions
async function GetPaypals() {
  const requestOptions = {
    method: "GET",
    headers: {
      "Content-Type": "application/json",
    },
  };

  try {
    const res = await fetch(`${apiUrl}/paypals`, requestOptions);
    if (res.status === 200) {
      return res.json();
    } else {
      return false;
    }
  } catch (error) {
    console.error("Error fetching PayPal accounts:", error);
    return false;
  }
}

async function GetPaypalById(id: number | undefined) {
  const requestOptions = {
    method: "GET",
    headers: {
      "Content-Type": "application/json",
    },
  };

  try {
    const res = await fetch(`${apiUrl}/paypals/${id}`, requestOptions);
    if (res.status === 200) {
      return res.json();
    } else {
      return false;
    }
  } catch (error) {
    console.error("Error fetching PayPal account by ID:", error);
    return false;
  }
}

async function CreatePaypal(data: PaypalInterface) {
  const requestOptions = {
    method: "POST",
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify(data),
  };

  try {
    const res = await fetch(`${apiUrl}/paypals`, requestOptions);
    if (res.status === 201) {
      return res.json();
    } else {
      return false;
    }
  } catch (error) {
    console.error("Error creating PayPal account:", error);
    return false;
  }
}

async function UpdatePaypal(data: PaypalInterface) {
  const requestOptions = {
    method: "PATCH",
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify(data),
  };

  try {
    const res = await fetch(`${apiUrl}/paypals/${data.PaypalID}`, requestOptions);
    if (res.status === 200) {
      return res.json();
    } else {
      return false;
    }
  } catch (error) {
    console.error("Error updating PayPal account:", error);
    return false;
  }
}

async function DeletePaypalByID(id: number | undefined) {
  const requestOptions = {
    method: "DELETE",
  };

  try {
    const res = await fetch(`${apiUrl}/paypals/${id}`, requestOptions);
    if (res.status === 200) {
      return true;
    } else {
      return false;
    }
  } catch (error) {
    console.error("Error deleting PayPal account:", error);
    return false;
  }
}

// PromptPay Functions
async function GetPromptPays() {
  const requestOptions = {
    method: "GET",
    headers: {
      "Content-Type": "application/json",
    },
  };

  try {
    const res = await fetch(`${apiUrl}/promptpays`, requestOptions);
    if (res.status === 200) {
      return res.json();
    } else {
      return false;
    }
  } catch (error) {
    console.error("Error fetching PromptPay accounts:", error);
    return false;
  }
}

async function GetPromptPayById(id: number | undefined) {
  const requestOptions = {
    method: "GET",
    headers: {
      "Content-Type": "application/json",
    },
  };

  try {
    const res = await fetch(`${apiUrl}/promptpays/${id}`, requestOptions);
    if (res.status === 200) {
      return res.json();
    } else {
      return false;
    }
  } catch (error) {
    console.error("Error fetching PromptPay account by ID:", error);
    return false;
  }
}

async function CreatePromptPay(data: PromptpayInterface) {
  const requestOptions = {
    method: "POST",
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify(data),
  };

  try {
    const res = await fetch(`${apiUrl}/promptpays`, requestOptions);
    if (res.status === 201) {
      return res.json();
    } else {
      return false;
    }
  } catch (error) {
    console.error("Error creating PromptPay account:", error);
    return false;
  }
}

async function UpdatePromptPay(data: PromptpayInterface) {
  const requestOptions = {
    method: "PATCH",
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify(data),
  };

  try {
    const res = await fetch(`${apiUrl}/promptpays/${data.PromptpayID}`, requestOptions);
    if (res.status === 200) {
      return res.json();
    } else {
      return false;
    }
  } catch (error) {
    console.error("Error updating PromptPay account:", error);
    return false;
  }
}

async function DeletePromptPayByID(id: number | undefined) {
  const requestOptions = {
    method: "DELETE",
  };

  try {
    const res = await fetch(`${apiUrl}/promptpays/${id}`, requestOptions);
    if (res.status === 200) {
      return true;
    } else {
      return false;
    }
  } catch (error) {
    console.error("Error deleting PromptPay account:", error);
    return false;
  }
}

export {
  GetCreditCards,
  GetCreditCardById,
  CreateCreditCard,
  UpdateCreditCard,
  DeleteCreditCardByID,
  GetPaypals,
  GetPaypalById,
  CreatePaypal,
  UpdatePaypal,
  DeletePaypalByID,
  GetPromptPays,
  GetPromptPayById,
  CreatePromptPay,
  UpdatePromptPay,
  DeletePromptPayByID,
};
