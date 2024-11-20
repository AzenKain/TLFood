package com.example.foodapp.adapter

import android.view.LayoutInflater
import android.view.View
import android.view.ViewGroup
import android.widget.ImageView
import android.widget.TextView
import androidx.recyclerview.widget.RecyclerView
import com.example.foodapp.R
import com.example.foodapp.data.Product

class ProductAdapter(var ls: List<Product>, private val onItemClick: (Product) -> Unit
): RecyclerView.Adapter<ProductAdapter.ProductViewHolder>(){

    inner class ProductViewHolder(item:View) :RecyclerView.ViewHolder(item), View.OnClickListener{
        val txtName:TextView = item.findViewById(R.id.product_name)
        val txtPrice:TextView = item.findViewById(R.id.product_price)
        val imageProduct:ImageView = item.findViewById(R.id.product_Image)
        val sales:TextView = item.findViewById(R.id.product_sale)
        init {
            item.setOnClickListener{
                onItemClick(ls[adapterPosition])
            }
        }
        override fun onClick(v: View?) {

        }
    }
    override fun onCreateViewHolder(parent: ViewGroup, viewType: Int): ProductViewHolder {
        val view = LayoutInflater.from(parent.context).inflate(R.layout.product_item, parent,false)
        return ProductViewHolder(view)
    }

    override fun getItemCount(): Int {
        return ls.size
    }
    override fun onBindViewHolder(holder: ProductViewHolder, position: Int) {
        holder.apply {
            txtName.text =ls[position].name
            val formatPrice = ls[position].price.formatPrice()
            val formatSales = ls[position].sales.formatSales()
            sales.text = formatSales
            txtPrice.text = formatPrice
            imageProduct.setImageResource(ls[position].image)
        }
    }
}
fun Int?.formatPrice(): String {
    return this?.let { "%,d".format(it).replace(',', '.') } ?: "0"
}
fun Int?.formatSales(): String {
    return when {
        this == null -> "0"
        this >= 1_000 -> String.format("%.1fk", this / 1_000.0)
        else -> this.toString()
    }
}
