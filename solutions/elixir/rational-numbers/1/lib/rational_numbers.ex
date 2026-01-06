defmodule RationalNumbers do
  @type rational :: {integer, integer}

  @doc """
  Add two rational numbers
  """
  @spec add(a :: rational, b :: rational) :: rational
  def add({a1, b1}, {a2, b2}) do
    # Sprowadzamy do wspólnego mianownika i dodajemy liczniki
    reduce({a1 * b2 + a2 * b1, b1 * b2})
  end

  @doc """
  Subtract two rational numbers
  """
  @spec subtract(a :: rational, b :: rational) :: rational
  def subtract({a1, b1}, {a2, b2}) do
    reduce({a1 * b2 - a2 * b1, b1 * b2})
  end

  @doc """
  Multiply two rational numbers
  """
  @spec multiply(a :: rational, b :: rational) :: rational
  def multiply({a1, b1}, {a2, b2}) do
    reduce({a1 * a2, b1 * b2})
  end

  @doc """
  Divide two rational numbers
  """
  @spec divide_by(num :: rational, den :: rational) :: rational
  def divide_by({a1, b1}, {a2, b2}) do
    # Mnożenie przez odwrotność drugiego ułamka
    reduce({a1 * b2, b1 * a2})
  end

  @doc """
  Absolute value of a rational number
  """
  @spec abs(a :: rational) :: rational
  def abs({n, d}) do
    # Wartość bezwzględna z licznika i mianownika + skrócenie
    reduce({Kernel.abs(n), Kernel.abs(d)})
  end

  @doc """
  Exponentiation of a rational number by an integer
  """
  @spec pow_rational(a :: rational, n :: integer) :: rational
  def pow_rational({num, den}, n) when n >= 0 do
    # Podnosimy licznik i mianownik do potęgi n
    reduce({Integer.pow(num, n), Integer.pow(den, n)})
  end

  def pow_rational({num, den}, n) when n < 0 do
    # Dla ujemnej potęgi odwracamy ułamek i podnosimy do potęgi dodatniej |n|
    m = Kernel.abs(n)
    reduce({Integer.pow(den, m), Integer.pow(num, m)})
  end

  @doc """
  Exponentiation of a real number by a rational number
  """
  @spec pow_real(x :: integer, n :: rational) :: float
  def pow_real(x, {num, den}) do
    # x do potęgi (num/den) -> pierwiastek stopnia den z (x do potęgi num)
    # W Elixirze używamy :math.pow(podstawa, wykładnik)
    :math.pow(x, num / den)
  end

  @doc """
  Reduce a rational number to its lowest terms
  """
  @spec reduce(a :: rational) :: rational
  def reduce({n, d}) do
    # Znajdź największy wspólny dzielnik (GCD)
    gcd = Integer.gcd(n, d)
    
    # Podziel licznik i mianownik przez GCD
    new_n = div(n, gcd)
    new_d = div(d, gcd)

    # Normalizacja znaku: Mianownik zawsze powinien być dodatni.
    # Jeśli mianownik jest ujemny, zmieniamy znaki obu liczb.
    if new_d < 0 do
      {-new_n, -new_d}
    else
      {new_n, new_d}
    end
  end
end