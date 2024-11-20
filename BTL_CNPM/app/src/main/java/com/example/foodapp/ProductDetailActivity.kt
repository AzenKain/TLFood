package com.example.foodapp

import android.os.Bundle
import androidx.activity.enableEdgeToEdge
import androidx.appcompat.app.AppCompatActivity
import androidx.core.view.ViewCompat
import androidx.core.view.WindowInsetsCompat
import com.example.foodapp.data.Product
import com.example.foodapp.databinding.ActivityProductDetailBinding

class ProductDetailActivity : AppCompatActivity() {
    private lateinit var binding :ActivityProductDetailBinding
    override fun onCreate(savedInstanceState: Bundle?) {
        super.onCreate(savedInstanceState)
        enableEdgeToEdge()
        setContentView(R.layout.activity_product_detail)
        ViewCompat.setOnApplyWindowInsetsListener(findViewById(R.id.main)) { v, insets ->
            val systemBars = insets.getInsets(WindowInsetsCompat.Type.systemBars())
            v.setPadding(systemBars.left, systemBars.top, systemBars.right, systemBars.bottom)
            insets
        }
        binding = ActivityProductDetailBinding.inflate(layoutInflater)
        setContentView(binding.root)

        val product: Product? = intent.getParcelableExtra("product")
        binding.productName.text = product?.name
        binding.productPrice.text = product?.price.toString()
        product?.image?.let { binding.productImage.setImageResource(it) }
        binding.productSale.text = product?.sales.toString()
        binding.productDescription.text = "Tung Hue"

        binding.btnAddToCart.setOnClickListener{
            addToCard()
        }
    }
}

private fun addToCard() {

}