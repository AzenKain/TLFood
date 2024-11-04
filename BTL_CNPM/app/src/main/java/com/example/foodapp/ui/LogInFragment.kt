package com.example.foodapp.ui

import android.os.Bundle
import android.os.Handler
import android.os.Looper
import android.text.Html
import androidx.fragment.app.Fragment
import android.view.LayoutInflater
import android.view.View
import android.view.ViewGroup
import androidx.navigation.fragment.findNavController
import com.example.foodapp.R
import com.example.foodapp.databinding.FragmentIntroductionBinding
import com.example.foodapp.databinding.FragmentLogInBinding


class LogInFragment : Fragment(R.layout.fragment_log_in) {
    private  var binding : FragmentLogInBinding? = null
    private val mBinding get() = binding!!
    override fun onCreateView(
        inflater: LayoutInflater,
        container: ViewGroup?,
        savedInstanceState: Bundle?
    ): View? {
        binding = FragmentLogInBinding.inflate(inflater, container, false)
        return mBinding.root
    }

    override fun onViewCreated(view: View, savedInstanceState: Bundle?) {
        super.onViewCreated(view, savedInstanceState)
        val nav = findNavController()
        val formattedText = "Don\\'t you have an account?<u><b> Sign up</b></u>"
        mBinding.tvDontAcc.text = Html.fromHtml(formattedText, Html.FROM_HTML_MODE_LEGACY)
        mBinding.tvDontAcc.setOnClickListener {
            nav.navigate(R.id.action_logInFragment_to_registerFragment)
        }
        mBinding.imgBackLogin.setOnClickListener {
            nav.popBackStack()
        }
        mBinding.btnLogIn.setOnClickListener {
            mBinding.progressBar.visibility = View.VISIBLE
        }
    }


}